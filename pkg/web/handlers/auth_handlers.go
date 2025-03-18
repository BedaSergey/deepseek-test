package handlers

import (
	"encoding/json"
	"net/http"
	"rent_alice/internal/service"
	"rent_alice/model"
)

type AuthHandlers struct {
	service *service.AuthService
}

func NewAuthHandlers(service *service.AuthService) *AuthHandlers {
	service.Logger.Info("Запускаем обработчики NewAuthHandlers")
	return &AuthHandlers{service: service}
}

// Handler регистрации пользователя
func (h *AuthHandlers) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		h.service.Logger.Error(err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if err := h.service.SignUp(r.Context(), model.User{
		Login:    request.Login,
		Password: request.Password,
	}); err != nil {
		h.service.Logger.Error(err)
		http.Error(w, "Registration failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

// Handler регистрации пользователя
func (h *AuthHandlers) SignInHandler(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		h.service.Logger.Error(err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, err := h.service.SignIn(r.Context(), request.Login, request.Password)
	if err != nil {
		h.service.Logger.Error(err)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	json_user, err := json.Marshal(user)
	// Генерация JWT токена (добавьте эту логику)
	// token := "newuser" + user.Login
	// token := generateToken(user)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"token": string(json_user),
	})
}
