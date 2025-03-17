package handlers

import (
	"net/http"
	"rent_alice/internal/service"
)

type AuthHandlers struct {
	service *service.AuthService
}

func NewAuthHandlers(service *service.AuthService) *AuthHandlers {
	service.Logger.Info("Запускаем обработчики NewAuthHandlers")
	return &AuthHandlers{service: service}
}

func (h *AuthHandlers) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	h.service.Logger.Info("Зашли в SignUpHandler")
}

func (h *AuthHandlers) SignInHandler(w http.ResponseWriter, r *http.Request) {
	h.service.Logger.Info("Зашли в SignInHandler")
}
