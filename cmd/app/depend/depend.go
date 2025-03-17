package depend

import (
	"rent_alice/internal/service"
	"rent_alice/pkg/web/handlers"
)

type Dependencies struct {
	DefaultService *service.DefaultService
	AuthService    *service.AuthService

	DefaultHandlers *handlers.Handlers
	AuthHandlers    *handlers.AuthHandlers
}
