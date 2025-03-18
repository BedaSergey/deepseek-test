package user_db

import (
	"context"
	"database/sql"
	"rent_alice/model"
	"rent_alice/pkg/database"
	"rent_alice/pkg/logger"
	"time"
)

type UserRepository struct {
	Database database.Database
	logger   logger.Logger
}

func NewUserRepository(logger logger.Logger, db database.Database) *UserRepository {
	return &UserRepository{
		logger:   logger,
		Database: db,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user model.User) error {
	var now = time.Now()
	qwery := "INSERT INTO users (login, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $4)"
	_, err := r.Database.Exec(ctx, qwery, user.Login, user.Password, now, now)
	if err != nil {
		r.logger.Error("r.Conn.Exec CreateUser failed:", err)
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByLogin(ctx context.Context, login string) (*model.User, error) {
	var user model.User
	query := `SELECT id, login, password_hash FROM users WHERE login = $1`
	err := r.Database.QueryRow(ctx, query, login).Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Пользователь не найден
		}
		r.logger.Error("Failed to get user by login:", err)
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CheckUser(ctx context.Context, user model.User) (bool, error) {
	var exist bool
	qwery := "SELECT EXISTS(SELECT 1 FROM users WHERE login = $1)"
	err := r.Database.QueryRow(ctx,
		qwery,
		user.Login,
	).Scan(&exist)
	if err != nil {
		r.logger.Error("r.Conn.QueryRow CheckUser failed:", err)
		return false, err
	}
	return true, nil
}
