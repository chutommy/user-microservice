package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

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

	CreateUser(ctx context.Context, username, password, firstName, lastName string, gender int16, email, phoneNumber string, birthDay time.Time) (repo.User, error)
	GetUserByID(ctx context.Context, id int64) (repo.User, error)
	GetUserByUsername(ctx context.Context, username string) (repo.User, error)
	GetUserByEmail(ctx context.Context, email string) (repo.User, error)
	UpdateUserUsername(ctx context.Context, id int64, username string) (repo.User, error)
	UpdateUserEmail(ctx context.Context, id int64, email string) (repo.User, error)
	UpdateUserPhoneNumber(ctx context.Context, id int64, phoneNumber string) (repo.User, error)
	UpdateUserPassword(ctx context.Context, id int64, password string) (repo.User, error)
	UpdateUserInfo(ctx context.Context, firstName, lastName, gender string, birthDay time.Time) (repo.User, error)
	DeleteUserSoft(ctx context.Context, id int64) error
	RecoverDeletedUser(ctx context.Context, id int64) (repo.User, error)
	DeleteUserPermanent(ctx context.Context, id int64) error
	VerifyPassword(ctx context.Context, id int64, password string) error
}

var (
	// ErrInternalDBError is returned if any unexpected database error occurs.
	ErrInternalDBError = errors.New("internal database error")
	// ErrNotFound is returned if no rows satisfy query search..
	ErrNotFound = errors.New("no found")
	// ErrUniqueGenderViolation is returned if gender with a duplicated title is requested to be added.
	ErrUniqueGenderViolation = errors.New("gender title must be unique")
	// ErrEmptyTitleField is returned if title field is empty.
	ErrEmptyTitleField = errors.New("title field is empty")
	// ErrEmptySearchKeys is returned if search keys are not provided.
	ErrEmptySearchKeys = errors.New("request search keys cannot be empty")
	// ErrEmptyID is returned if empty ID value is provided.
	ErrEmptyID = errors.New("id field cannot be of null value")
)

type basicUserService struct {
	repo repo.Querier
}

func (b *basicUserService) AddGender(ctx context.Context, title string) (repo.Gender, error) {
	if title == "" {
		return repo.Gender{}, ErrEmptyTitleField
	}

	// insert
	g, err := b.repo.CreateGender(ctx, title)
	if err != nil {
		if err, ok := err.(pq.Error); ok {
			switch err.Code {
			case "23505": // unique key violation
				return repo.Gender{}, ErrUniqueGenderViolation
			}
		}

		return repo.Gender{}, fmt.Errorf("could not create a new gender: %w", multierr.Append(ErrInternalDBError, err))
	}

	return g, nil
}
func (b *basicUserService) GetGender(ctx context.Context, id int16) (repo.Gender, error) {
	if id == 0 {
		return repo.Gender{}, ErrEmptySearchKeys
	}

	// select
	g, err := b.repo.GetGender(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return repo.Gender{}, ErrNotFound
		}

		return repo.Gender{}, fmt.Errorf("could not select a gender: %w", multierr.Append(ErrInternalDBError, err))
	}

	return g, nil
}
func (b *basicUserService) ListGenders(ctx context.Context) ([]repo.Gender, error) {

	genders, err := b.repo.ListGenders(ctx)
	if err != nil {
		return []repo.Gender{}, fmt.Errorf("could not list genders: %w", multierr.Append(ErrInternalDBError, err))
	}

	return genders, nil
}
func (b *basicUserService) RemoveGender(ctx context.Context, id int16) error {
	if id == 0 {
		return ErrEmptyID
	}

	if err := b.repo.DeleteGender(ctx, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNotFound
		}

		return fmt.Errorf("could not delete gender: %w", multierr.Append(ErrInternalDBError, err))
	}

	return nil
}
func (b *basicUserService) CreateUser(ctx context.Context, username string, password string, firstName string, lastName string, gender int16, email string, phoneNumber string, birthDay time.Time) (r0 repo.User, e1 error) {
	// TODO implement the business logic of CreateUser
	return r0, e1
}
func (b *basicUserService) GetUserByID(ctx context.Context, id int64) (r0 repo.User, e1 error) {
	// TODO implement the business logic of GetUserByID
	return r0, e1
}
func (b *basicUserService) GetUserByUsername(ctx context.Context, username string) (r0 repo.User, e1 error) {
	// TODO implement the business logic of GetUserByUsername
	return r0, e1
}
func (b *basicUserService) GetUserByEmail(ctx context.Context, email string) (r0 repo.User, e1 error) {
	// TODO implement the business logic of GetUserByEmail
	return r0, e1
}
func (b *basicUserService) UpdateUserUsername(ctx context.Context, id int64, username string) (r0 repo.User, e1 error) {
	// TODO implement the business logic of UpdateUserUsername
	return r0, e1
}
func (b *basicUserService) UpdateUserEmail(ctx context.Context, id int64, email string) (r0 repo.User, e1 error) {
	// TODO implement the business logic of UpdateUserEmail
	return r0, e1
}
func (b *basicUserService) UpdateUserPhoneNumber(ctx context.Context, id int64, phoneNumber string) (r0 repo.User, e1 error) {
	// TODO implement the business logic of UpdateUserPhoneNumber
	return r0, e1
}
func (b *basicUserService) UpdateUserPassword(ctx context.Context, id int64, password string) (r0 repo.User, e1 error) {
	// TODO implement the business logic of UpdateUserPassword
	return r0, e1
}
func (b *basicUserService) UpdateUserInfo(ctx context.Context, firstName string, lastName string, gender string, birthDay time.Time) (r0 repo.User, e1 error) {
	// TODO implement the business logic of UpdateUserInfo
	return r0, e1
}
func (b *basicUserService) DeleteUserSoft(ctx context.Context, id int64) (e0 error) {
	// TODO implement the business logic of DeleteUserSoft
	return e0
}
func (b *basicUserService) RecoverDeletedUser(ctx context.Context, id int64) (r0 repo.User, e1 error) {
	// TODO implement the business logic of RecoverDeletedUser
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
	return &basicUserService{
		repo: repo,
	}
}

// New returns a UserService with all of the expected middleware wired in.
func New(repo repo.Querier, middleware []Middleware) UserService {
	var svc = NewBasicUserService(repo)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
