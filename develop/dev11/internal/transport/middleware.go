package transport

import (
	"log/slog"
	"net/http"
	"time"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(lrw, r)

		duration := time.Since(start)
		if lrw.statusCode != http.StatusOK {
			slog.Error("Error handling request",
				slog.String("method", r.Method),
				slog.String("url", r.URL.String()),
				slog.Int("status", lrw.statusCode),
				slog.Duration("duration", duration),
			)
		} else {
			slog.Info("Handled request",
				slog.String("method", r.Method),
				slog.String("url", r.URL.String()),
				slog.Int("status", lrw.statusCode),
				slog.Duration("duration", duration),
			)
		}
	})
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *loggingResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *loggingResponseWriter) Write(b []byte) (int, error) {
	return w.ResponseWriter.Write(b)
}
