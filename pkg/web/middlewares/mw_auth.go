package middlewares

import (
	"net/http"
	"rent_alice/internal/service"
)

func IsAuth(s *service.AuthService) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Необходимо выполнять проверку сессии
			s.Logger.Info("Проходили через IsAuth")
			next.ServeHTTP(w, r)
		})
	}
}

// func IsAuth(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Необходимо выполнять проверку сессии
// 		s.Logger.Info("Проходили через IsAuth")
// 		next.ServeHTTP(w, r)
// 	})
// }
