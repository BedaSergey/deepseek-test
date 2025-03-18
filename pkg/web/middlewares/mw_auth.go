package middlewares

import (
	"encoding/json"
	"net/http"
	"rent_alice/internal/service"
)

func NewIsAuthMiddleware(s *service.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var request struct {
				Login string `json:"login"`
			}
			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
				s.Logger.Error(err)
				http.Error(w, "json.NewDecoder Invalid request", http.StatusBadRequest)
				return
			}
			exist, err := s.UserExist(r.Context(), request.Login)
			if err != nil {
				s.Logger.Error(err)
				http.Error(w, "s.UserExist Invalid request", http.StatusBadRequest)
				return
			}
			if exist {
				s.Logger.Info("Пользователь был", exist)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
