package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

// LoggingMiddleware logs the details of the incoming HTTP request and the time taken to process it.
func LoggingMiddleware(next http.Handler, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		logger.Info(fmt.Sprintf("%s %s %s", r.Method, r.URL.String(), time.Since(startTime).String()))
	})
}
