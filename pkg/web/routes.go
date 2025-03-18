package web

import (
	"rent_alice/cmd/app/depend"

	"github.com/gorilla/mux"
)

func NewRouter(deps depend.Dependencies) *mux.Router {

	r := mux.NewRouter()

	// Подроутер для аутентификации
	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/signup", deps.AuthHandlers.SignUpHandler).Methods("POST")
	authRouter.HandleFunc("/signin", deps.AuthHandlers.SignInHandler).Methods("POST")
	// Используемые Middlewares
	authRouter.Use(deps.Middlewares.Logging)

	// Подроутер для остальных маршрутов
	defaultRouter := r.PathPrefix("/").Subrouter()
	defaultRouter.HandleFunc("/", deps.DefaultHandlers.HomeHandler)
	// Используемые Middlewares
	defaultRouter.Use(deps.Middlewares.Logging, deps.Middlewares.IsAuth)

	return r
}
