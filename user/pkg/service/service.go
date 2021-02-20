package service

import (
	"context"
	"time"

	"user/pkg/repo"
)

// UserService describes the service.
type UserService interface {
	AddGender(ctx context.Context, title string) (repo.Gender, error)
	GetGender(ctx context.Context, id int16) (repo.Gender, error)
	ListGenders(ctx context.Context) ([]repo.Gender, error)
	RemoveGender(ctx context.Context, id int16) error

	CreateUser(ctx context.Context, password, firstName, lastName string, gender int16, email, phoneNumber string, birthDay time.Time) (repo.User, error)
	GetUserByID(ctx context.Context, id int64) (repo.User, error)
	GetUserByEmail(ctx context.Context, email string) (repo.User, error)
	UpdateUserEmail(ctx context.Context, id int64, email string) (repo.User, error)
	UpdateUserPhoneNumber(ctx context.Context, id int64, phoneNumber string) (repo.User, error)
	UpdateUserPassword(ctx context.Context, id int64, password string) (repo.User, error)
	UpdateUserInfo(ctx context.Context, firstName, lastName, gender string, birthDay time.Time) (repo.User, error)
	DeleteUserSoft(ctx context.Context, id int64) error
	RecoverDeletedUser(ctx context.Context, id int64) (repo.User, error)
	DeleteUserPermanent(ctx context.Context, id int64) error
	VerifyPassword(ctx context.Context, id int64, password string) error
}
