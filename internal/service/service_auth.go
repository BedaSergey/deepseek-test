package service

import (
	"rent_alice/pkg/database"
	"rent_alice/pkg/logger"
)

type AuthService struct {
	Logger   logger.Logger
	Database database.Database
}

func NewAuthService(logger logger.Logger, database database.Database) *AuthService {
	return &AuthService{
		Logger:   logger,
		Database: database}
}
