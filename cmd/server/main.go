package main

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"text/template"
	"time"

	"github.com/ctbur/ctbur.me/internal/fragments"
	"github.com/ctbur/ctbur.me/internal/log"
)

func main() {
	tmpl, err := template.ParseGlob("templates/*.html")
	if err != nil {
		slog.Error("Failed to load templates", slog.Any("error", err))
		return
	}

	mux := http.NewServeMux()

	staticFileServer := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFileServer))

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		log := log.FromContext(r.Context())
		if RenderPage(*tmpl, "page_home", nil, w); err != nil {
			log.Error("Failed to render page", slog.Any("error", err))
		}
	})

	fragments, err := fragments.LoadFragments("content/fragments.toml")
	if err != nil {
		slog.Error("Failed to load fragments", slog.Any("error", err))
		return
	}

	mux.HandleFunc("GET /fragments", func(w http.ResponseWriter, r *http.Request) {
		log := log.FromContext(r.Context())
		if RenderPage(*tmpl, "page_fragments", fragments, w); err != nil {
			log.Error("Failed to render page", slog.Any("error", err))
		}
	})

	handler := log.Middleware(mux)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{Addr: fmt.Sprintf(":%s", port), Handler: handler}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				slog.Error("Fatal error", slog.Any("error", err))
			}
		}
	}()

	// Wait for SIGINT
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Fatal error", slog.Any("error", err))
	}
}

func RenderPage(tmpl template.Template, pageTemplate string, data any, w http.ResponseWriter) error {
	var b bytes.Buffer
	err := tmpl.ExecuteTemplate(&b, pageTemplate, data)
	if err != nil {
		// TODO: error page
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	b.WriteTo(w)
	return nil
}
