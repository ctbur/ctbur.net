package main

import (
	"bytes"
	"context"
	"errors"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ctbur/ctbur.me/internal/log"
	"github.com/ctbur/ctbur.me/internal/til"
)

func loadTemplates() (*template.Template, error) {
	tmpl := template.New("main")
	tmpl.Funcs(template.FuncMap{
		"dict": dict,
	})

	return tmpl.ParseGlob("templates/*.html")
}

func dict(values ...any) (map[string]any, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("invalid dict call: number of arguments must be even")
	}

	dict := make(map[string]any, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}

func main() {
	tmpl, err := loadTemplates()
	if err != nil {
		slog.Error("Failed to load templates", slog.Any("error", err))
		return
	}

	mux := http.NewServeMux()

	staticFileServer := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFileServer))

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		log := log.FromContext(r.Context())

		var b bytes.Buffer
		err := tmpl.ExecuteTemplate(&b, "page_index", nil)
		if err != nil {
			log.Error("Failed to render page", slog.Any("error", err))
			// TODO: error page
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		b.WriteTo(w)
	})

	tils, err := til.LoadTils("content/tils.toml")
	slog.Info("Tils", slog.Int("num", len(tils)))
	if err != nil {
		slog.Error("Failed to load TILs", slog.Any("error", err))
		return
	}

	mux.HandleFunc("GET /til", func(w http.ResponseWriter, r *http.Request) {
		log := log.FromContext(r.Context())

		var b bytes.Buffer
		err := tmpl.ExecuteTemplate(&b, "page_til", tils)
		if err != nil {
			log.Error("Failed to render page", slog.Any("error", err))
			// TODO: error page
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		b.WriteTo(w)
	})

	handler := log.Middleware(mux)
	server := &http.Server{Addr: ":8080", Handler: handler}

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
