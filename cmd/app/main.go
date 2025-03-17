package main

import (
	"net/http"
	"os"
	"rent_alice/cmd/app/depend"
	"rent_alice/internal/service"
	"rent_alice/pkg/database"
	"rent_alice/pkg/logger"
	"rent_alice/pkg/web"
	"rent_alice/pkg/web/handlers"
)

func main() {
	logger := logger.NewLogrusLogger()

	db := database.NewPostgresDatabase(logger)

	defaultService := service.NewDefaultService(logger, db)
	serviceAuth := service.NewAuthService(logger, db)

	defaultHandlers := handlers.NewHandlers(defaultService)
	authHandlers := handlers.NewAuthHandlers(serviceAuth)

	deps := &depend.Dependencies{
		DefaultService:  defaultService,
		AuthService:     serviceAuth,
		DefaultHandlers: defaultHandlers,
		AuthHandlers:    authHandlers}

	r := web.NewRouter(*deps)

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080" // Порт по умолчанию
	}

	logger.Info("Starting server on :" + port)
	if err := http.ListenAndServe(port, r); err != nil {
		logger.Error("Server failed:", err)
	}
}
