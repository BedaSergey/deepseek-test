package main

import (
	"net/http"
	"os"
	"rent_alice/cmd/app/depend"
	"rent_alice/internal/service"
	"rent_alice/pkg/database"
	"rent_alice/pkg/database/user_db"
	"rent_alice/pkg/logger"
	"rent_alice/pkg/web"
	"rent_alice/pkg/web/handlers"
	"rent_alice/pkg/web/middlewares"
)

func main() {
	logger := logger.NewLogrusLogger()

	db := database.NewPostgresDatabase(logger)
	defer db.Close()

	// Загрузка репозиториев, которые используются для взаимодейтсвия с базой данных из сервисов
	userRepo := user_db.NewUserRepository(logger, db)

	// Загрузка сервисов, которые взаимодействуют с репозиториями из хэндлеров
	defaultService := service.NewDefaultService(logger, db) // Можно убрать db, потому что сервис их не использует. Но если не убирать, тогда заменить на кастомный
	serviceAuth := service.NewAuthService(logger, userRepo)

	// Загрузка хэндлеров
	defaultHandlers := handlers.NewHandlers(defaultService)
	authHandlers := handlers.NewAuthHandlers(serviceAuth)

	// Загрузка middlewares
	authMiddleware := middlewares.NewIsAuthMiddleware(serviceAuth)
	loggingMiddleware := middlewares.NewLoggingMiddleware(logger)

	deps := &depend.Dependencies{
		DefaultHandlers: defaultHandlers,
		AuthHandlers:    authHandlers,
		Middlewares: depend.Middlewares{
			IsAuth:  authMiddleware,
			Logging: loggingMiddleware,
		},
	}

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
