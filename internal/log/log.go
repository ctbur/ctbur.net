package log

import (
	"context"
	"log/slog"
	"net/http"
	"time"
)

type loggerContextKeyType string

const loggerContextKey loggerContextKeyType = "requestLogger"

// logWriter wraps an http.ResponseWriter to capture info about the result
type logWriter struct {
	http.ResponseWriter
	statusCode int
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Create a context with a request logger
		requestLogger := slog.Default().With(
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
		)
		ctx := context.WithValue(r.Context(), loggerContextKey, requestLogger)

		lw := &logWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK, // Default status code if WriteHeader is not called
		}
		next.ServeHTTP(lw, r.WithContext(ctx))

		// Log request and duration
		duration := time.Since(startTime)
		requestLogger.Info(
			"Request",
			slog.Int64("duration_ms", duration.Milliseconds()),
		)
	})
}

func FromContext(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(loggerContextKey).(*slog.Logger); ok {
		return logger
	}

	// Return a default logger or handle the error if the logger is not found.
	// This might happen if a handler is called directly without the middleware.
	return slog.Default()
}
