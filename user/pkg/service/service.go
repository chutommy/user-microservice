package service

import (
	"context"

	"user/pkg/repo"
)

// UserService describes the service.
type UserService interface {
	AddGender(ctx context.Context, title string) (repo.Gender, error)
	GetGender(ctx context.Context, id int16) (repo.Gender, error)
	ListGenders(ctx context.Context) ([]repo.Gender, error)
	RemoveGender(ctx context.Context, id int16) error

	CreateUser(ctx context.Context, user repo.User) (repo.User, error)
	GetUserByID(ctx context.Context, id int64) (repo.User, error)
	GetUserByEmail(ctx context.Context, email string) (repo.User, error)
	UpdateUserEmail(ctx context.Context, id int64, email string) (repo.User, error)
	UpdateUserPassword(ctx context.Context, id int64, password string) (repo.User, error)
	UpdateUserInfo(ctx context.Context, id int64, user repo.User) (repo.User, error)
	DeleteUserSoft(ctx context.Context, id int64) error
	RecoverUser(ctx context.Context, id int64) (repo.User, error)
	DeleteUserPermanent(ctx context.Context, id int64) error

	VerifyPassword(ctx context.Context, id int64, password string) error
}
