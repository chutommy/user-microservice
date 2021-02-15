package service

import (
	"context"
	"time"
)

// UserService describes the service.
type UserService interface {
	AddGender(ctx context.Context, title string) (int64, error)
	GetGender(ctx context.Context, id int64, title string) (Gender, error)
	ListGenders(ctx context.Context) ([]Gender, error)
	RemoveGender(ctx context.Context, title string) error

	CreateUser(ctx context.Context, username, password, firstName, lastName, gender, email, phoneNumber string, birthDay time.Time) (User, error)
	GetUserByID(ctx context.Context, id int64) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	UpdateUserUsername(ctx context.Context, id int64, username string) (User, error)
	UpdateUserEmail(ctx context.Context, id int64, email string) (User, error)
	UpdateUserPhoneNumber(ctx context.Context, id int64, phoneNumber string) (User, error)
	UpdateUserPassword(ctx context.Context, id int64, password string) (User, error)
	UpdateUserInfo(ctx context.Context, firstName, lastName, gender string, birthDay time.Time) (User, error)
	DeleteUserSoft(ctx context.Context, id int64) error
	RecoverDeletedUser(ctx context.Context, id int64) (User, error)
	DeleteUserPermanent(ctx context.Context, id int64) error
	VerifyPassword(ctx context.Context, id int64, password string) error
}
