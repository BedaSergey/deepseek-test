package web

import (
	"net/http"
	"rent_alice/cmd/app/depend"
	"rent_alice/pkg/web/middlewares"

	"github.com/gorilla/mux"
)

func NewRouter(deps depend.Dependencies) *mux.Router {

	r := mux.NewRouter()

	// Подроутер для аутентификации
	authRouter := r.PathPrefix("/auth").Subrouter()
	authRouter.Handle("/signup",
		middlewares.IsAuth(deps.AuthService)(
			http.HandlerFunc(deps.AuthHandlers.SignUpHandler)))
	authRouter.Handle("/signin",
		middlewares.IsAuth(deps.AuthService)(
			http.HandlerFunc(deps.AuthHandlers.SignInHandler)))

	// Подроутер для остальных маршрутов
	defaultRouter := r.PathPrefix("/").Subrouter()
	defaultRouter.HandleFunc("/", deps.DefaultHandlers.HomeHandler)

	return r

}

// func NewAuthRouter(r *mux.Router, handlers *handlers.AuthHandlers, service *service.AuthService) *mux.Router {

// 	// Роуты доступные с авторизацией
// 	r.Handle("/signup",
// 		middlewares.IsAuth(service)(
// 			http.HandlerFunc(handlers.SignUpHandler)))
// 	r.Handle("/signin",
// 		middlewares.IsAuth(service)(
// 			http.HandlerFunc(handlers.SignInHandler)))

// 	return r

// }
