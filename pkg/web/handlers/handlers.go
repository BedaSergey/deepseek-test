package handlers

import (
	"net/http"
	"rent_alice/internal/service"
)

type Handlers struct {
	service *service.DefaultService
}

func NewHandlers(service *service.DefaultService) *Handlers {
	service.Logger.Info("Запускаем обработчики")
	return &Handlers{service: service}
}

func (h *Handlers) HomeHandler(w http.ResponseWriter, r *http.Request) {
	h.service.Logger.Info("Зашли в HomeHandlers")
}
