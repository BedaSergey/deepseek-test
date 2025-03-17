package service

import (
	"rent_alice/pkg/database"
	"rent_alice/pkg/logger"
)

type DefaultService struct {
	Logger   logger.Logger
	Database database.Database
}

func NewDefaultService(logger logger.Logger, database database.Database) *DefaultService {
	return &DefaultService{
		Logger:   logger,
		Database: database}
}
