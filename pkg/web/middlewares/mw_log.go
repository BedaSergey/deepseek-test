package middlewares

import (
	"net/http"
	"rent_alice/pkg/logger"
)

func NewLoggingMiddleware(logger logger.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("Request: ", r.Method, " ", r.URL.Path, " ", r.Body)
			next.ServeHTTP(w, r)
		})
	}
}
