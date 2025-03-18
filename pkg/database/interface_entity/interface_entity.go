package interface_entity

import (
	"context"
	"rent_alice/model"
)

type Repository interface {
	Create(ctx context.Context, entity interface{}) error
	Update(ctx context.Context, entity interface{}) error
	Delete(ctx context.Context, id string) error
	FindByID(ctx context.Context, id string) (interface{}, error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, user model.User) error
	GetUserByLogin(ctx context.Context, login string) (*model.User, error)
	CheckUser(ctx context.Context, user model.User) (bool, error)
}
