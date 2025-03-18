package service

import (
	"context"
	"errors"
	"rent_alice/model"
	"rent_alice/pkg/database/interface_entity"
	"rent_alice/pkg/logger"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Logger logger.Logger
	// Database database.Database
	userRepo interface_entity.UserRepository
}

func NewAuthService(logger logger.Logger, userRepo interface_entity.UserRepository) *AuthService {
	return &AuthService{
		Logger:   logger,
		userRepo: userRepo}
}

// Сервис регистрации пользователя
func (s *AuthService) SignUp(ctx context.Context, user model.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		s.Logger.Error("Password hashing failed:", err)
		return err
	}

	user.Password = string(hashedPassword)

	err = s.userRepo.CreateUser(ctx, user)
	if err != nil {
		s.Logger.Error("SignUp Service failed:", err)
		return err
	}

	return nil
}

// Сервис авторизации пользователя
func (s *AuthService) SignIn(ctx context.Context, login, password string) (*model.User, error) {
	var user *model.User

	user, err := s.userRepo.GetUserByLogin(ctx, login)
	if err != nil {
		s.Logger.Error("s.userRepo.CheckUser SignIn failed:", err)
		return nil, errors.New("invalid credentials")
	}

	// Проверяем пароль
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		s.Logger.Error("bcrypt.CompareHashAndPassword SignIn failed:", err)
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *AuthService) UserExist(ctx context.Context, login string) (bool, error) {
	user, err := s.userRepo.GetUserByLogin(ctx, login)
	if err != nil {
		s.Logger.Error("s.userRepo.CheckUser SignIn failed:", err)
		return false, errors.New("invalid credentials")
	}
	if user == nil {
		return false, nil
	}
	return true, nil
}
