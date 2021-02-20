package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/lib/pq"
	"go.uber.org/multierr"

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

// basicUserService implements UserService interface.
type basicUserService struct {
	repo repo.Querier
}

var (
	// ErrInternalDBError is returned if any unexpected database error occurs.
	ErrInternalDBError = errors.New("internal database error")
	// ErrMissingRequestField is returned if any of the mandatory parameter field is left empty.
	ErrMissingRequestField = errors.New("request are incomplete")
	// ErrDuplicatedValue is returned whenever any duplication of a value that must be unique appears.
	ErrDuplicatedValue = errors.New("given value is already in use")
)

func (b *basicUserService) AddGender(ctx context.Context, title string) (repo.Gender, error) {
	if title == "" {
		return repo.Gender{}, fmt.Errorf("%w: missing title", ErrMissingRequestField)
	}

	g, err := b.repo.CreateGender(ctx, title)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Code {
			case "23505":
				return repo.Gender{}, fmt.Errorf("%w: duplicated 'title'", ErrDuplicatedValue)
			}

			return repo.Gender{}, multierr.Append(
				ErrInternalDBError,
				fmt.Errorf("failed to create gender: %w", err),
			)
		}
	}

	return g, err
}
func (b *basicUserService) GetGender(ctx context.Context, id int16) (r0 repo.Gender, e1 error) {
	// TODO implement the business logic of GetGender
	return r0, e1
}
func (b *basicUserService) ListGenders(ctx context.Context) (r0 []repo.Gender, e1 error) {
	// TODO implement the business logic of ListGenders
	return r0, e1
}
func (b *basicUserService) RemoveGender(ctx context.Context, id int16) (e0 error) {
	// TODO implement the business logic of RemoveGender
	return e0
}
func (b *basicUserService) CreateUser(ctx context.Context, user repo.User) (r0 repo.User, e1 error) {
	// TODO implement the business logic of CreateUser
	return r0, e1
}
func (b *basicUserService) GetUserByID(ctx context.Context, id int64) (r0 repo.User, e1 error) {
	// TODO implement the business logic of GetUserByID
	return r0, e1
}
func (b *basicUserService) GetUserByEmail(ctx context.Context, email string) (r0 repo.User, e1 error) {
	// TODO implement the business logic of GetUserByEmail
	return r0, e1
}
func (b *basicUserService) UpdateUserEmail(ctx context.Context, id int64, email string) (r0 repo.User, e1 error) {
	// TODO implement the business logic of UpdateUserEmail
	return r0, e1
}
func (b *basicUserService) UpdateUserPassword(ctx context.Context, id int64, password string) (r0 repo.User, e1 error) {
	// TODO implement the business logic of UpdateUserPassword
	return r0, e1
}
func (b *basicUserService) UpdateUserInfo(ctx context.Context, id int64, user repo.User) (r0 repo.User, e1 error) {
	// TODO implement the business logic of UpdateUserInfo
	return r0, e1
}
func (b *basicUserService) DeleteUserSoft(ctx context.Context, id int64) (e0 error) {
	// TODO implement the business logic of DeleteUserSoft
	return e0
}
func (b *basicUserService) RecoverUser(ctx context.Context, id int64) (r0 repo.User, e1 error) {
	// TODO implement the business logic of RecoverUser
	return r0, e1
}
func (b *basicUserService) DeleteUserPermanent(ctx context.Context, id int64) (e0 error) {
	// TODO implement the business logic of DeleteUserPermanent
	return e0
}
func (b *basicUserService) VerifyPassword(ctx context.Context, id int64, password string) (e0 error) {
	// TODO implement the business logic of VerifyPassword
	return e0
}

// NewBasicUserService returns a naive, stateless implementation of UserService.
func NewBasicUserService(repo repo.Querier) UserService {
	return &basicUserService{repo}
}

// New returns a UserService with all of the expected middleware wired in.
func New(repo repo.Querier, middleware []Middleware) UserService {
	var svc UserService = NewBasicUserService(repo)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
