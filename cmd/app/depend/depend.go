package depend

import (
	"net/http"
	"rent_alice/pkg/web/handlers"
)

type Dependencies struct {
	DefaultHandlers *handlers.Handlers
	AuthHandlers    *handlers.AuthHandlers

	Middlewares Middlewares
}

type Middlewares struct {
	IsAuth  func(http.Handler) http.Handler
	Logging func(http.Handler) http.Handler
}
