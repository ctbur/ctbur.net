package main

import (
	"context"
	"html/template"
	"io/fs"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"time"

	log "github.com/ctbur/ctbur.me/internal"
)

func loadTemplates() (*template.Template, error) {
	var tmplFiles []string
	filepath.Walk("templates", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".html") {
			tmplFiles = append(tmplFiles, path)
		}

		return nil
	})

	return template.ParseFiles(tmplFiles...)
}

func main() {
	tmpl, err := loadTemplates()
	if err != nil {
		slog.Error("Failed to load templates", slog.Any("error", err))
		return
	}

	mux := http.NewServeMux()

	staticFileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", staticFileServer))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log := log.FromContext(r.Context())
		err := tmpl.ExecuteTemplate(w, "page_index", nil)
		if err != nil {
			log.Error("Failed to render page", slog.Any("error", err))
		}
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
