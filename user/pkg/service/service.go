package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
	"go.uber.org/multierr"
	"golang.org/x/crypto/bcrypt"

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
	// ErrPasswordHashing is returned whenever unexpected error is arisen from hashing passwords.
	ErrPasswordHashing = errors.New("password hashing error")
	// ErrInternalServerError is returned if any unexpected server error occurs.
	ErrInternalServerError = errors.New("internal server error")
	// ErrMissingRequestField is returned if any of the mandatory parameter field is left empty.
	ErrMissingRequestField = errors.New("request are incomplete")
	// ErrDuplicatedValue is returned whenever any duplication of a value that must be unique appears.
	ErrDuplicatedValue = errors.New("given value is already in use")
	// ErrNotFound is returned if searched value is not found.
	ErrNotFound = errors.New("record not found")
	// ErrWrongPassword is returned if the given password doesn't match to the
	// password that was stored with the specific ID.
	ErrWrongPassword = errors.New("passwords don't match")
)

func (b *basicUserService) AddGender(ctx context.Context, title string) (repo.Gender, error) {
	if title == "" {
		return repo.Gender{}, fmt.Errorf("%w: missing 'title'", ErrMissingRequestField)
	}

	// create
	g, err := b.repo.CreateGender(ctx, title)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Code {
			case "23505":
				return repo.Gender{}, fmt.Errorf("%w: duplicated 'title'", ErrDuplicatedValue)
			}
		}

		return repo.Gender{}, multierr.Append(
			ErrInternalServerError,
			fmt.Errorf("failed to create gender: %w", err),
		)
	}

	return g, err
}
func (b *basicUserService) GetGender(ctx context.Context, id int16) (repo.Gender, error) {
	if id == 0 {
		return repo.Gender{}, fmt.Errorf("%w: missing 'id'", ErrMissingRequestField)
	}

	// retrieve
	g, err := b.repo.GetGender(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return repo.Gender{}, fmt.Errorf("%w: no gender with id '%d'", ErrNotFound, id)
		}

		return repo.Gender{}, multierr.Append(
			ErrInternalServerError,
			fmt.Errorf("failed to retrieve gender: %w", err),
		)
	}

	return g, nil
}

func (b *basicUserService) ListGenders(ctx context.Context) ([]repo.Gender, error) {
	// retrieve
	gs, err := b.repo.ListGenders(ctx)
	if err != nil {
		return []repo.Gender{}, multierr.Append(
			ErrInternalServerError,
			fmt.Errorf("failed to retrieve genders: %w", err),
		)
	}

	return gs, nil
}
func (b *basicUserService) RemoveGender(ctx context.Context, id int16) error {
	if id == 0 {
		return fmt.Errorf("%w: missing 'id'", ErrMissingRequestField)
	}

	// delete
	err := b.repo.DeleteGender(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("%w: no gender with id '%d'", ErrNotFound, id)
		}

		return multierr.Append(ErrInternalServerError, fmt.Errorf("failed to delte gender: %w", err))
	}

	return nil
}

func (b *basicUserService) CreateUser(ctx context.Context, user repo.User) (repo.User, error) {
	switch {
	case user.Email == "":
		return repo.User{}, fmt.Errorf("%w: missing 'email'", ErrMissingRequestField)
	case user.HashedPassword == "":
		return repo.User{}, fmt.Errorf("%w: missing 'hashedPassword'", ErrMissingRequestField)
	case user.FirstName == "":
		return repo.User{}, fmt.Errorf("%w: missing 'firstName'", ErrMissingRequestField)
	case user.LastName == "":
		return repo.User{}, fmt.Errorf("%w: missing 'lastName'", ErrMissingRequestField)
	case user.Gender == 0:
		return repo.User{}, fmt.Errorf("%w: missing 'gender'", ErrMissingRequestField)
	}

	bp, err := bcrypt.GenerateFromPassword([]byte(user.HashedPassword), bcrypt.DefaultCost)
	if err != nil {
		return repo.User{}, multierr.Append(ErrInternalServerError, ErrPasswordHashing)
	}
	user.HashedPassword = string(bp)

	newUser, err := b.repo.CreateUser(ctx, repo.CreateUserParams{
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		BirthDay:       user.BirthDay,
		Gender:         user.Gender,
		PhoneNumber:    user.PhoneNumber,
	})
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Code {
			case "23505":
				return repo.User{}, multierr.Append(ErrDuplicatedValue, pqErr)
			}
		}

		return repo.User{}, multierr.Append(
			ErrInternalServerError,
			fmt.Errorf("failed to create user: %w", err),
		)
	}

	return newUser, nil
}
func (b *basicUserService) GetUserByID(ctx context.Context, id int64) (repo.User, error) {
	if id == 0 {
		return repo.User{}, fmt.Errorf("%w: missing 'id'", ErrMissingRequestField)
	}

	u, err := b.repo.GetUserByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return repo.User{}, fmt.Errorf("%w: no user with 'id': %d", ErrNotFound, id)
		}

		return repo.User{}, multierr.Append(ErrInternalServerError, err)
	}

	return u, nil
}
func (b *basicUserService) GetUserByEmail(ctx context.Context, email string) (repo.User, error) {
	if email == "" {
		return repo.User{}, fmt.Errorf("%w: missing 'email'", ErrMissingRequestField)
	}

	u, err := b.repo.GetUserByEmail(ctx, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return repo.User{}, fmt.Errorf("%w: no user with 'email': %s", ErrNotFound, email)
		}

		return repo.User{}, multierr.Append(ErrInternalServerError, err)
	}

	return u, nil
}
func (b *basicUserService) UpdateUserEmail(ctx context.Context, id int64, email string) (repo.User, error) {
	switch {
	case id == 0:
		return repo.User{}, fmt.Errorf("%w: missing 'id'", ErrMissingRequestField)
	case email == "":
		return repo.User{}, fmt.Errorf("%w: missing 'email'", ErrMissingRequestField)
	}

	u, err := b.repo.UpdateUserEmail(ctx, repo.UpdateUserEmailParams{
		ID:    id,
		Email: email,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return repo.User{}, fmt.Errorf("%w: no user with 'id': %d", ErrNotFound, id)
		}

		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Code {
			case "23505":
				return repo.User{}, multierr.Append(ErrDuplicatedValue, pqErr)
			}
		}

		return repo.User{}, multierr.Append(
			ErrInternalServerError,
			fmt.Errorf("failed to update user's 'email': %w", err),
		)
	}

	return u, nil
}

func (b *basicUserService) UpdateUserPassword(ctx context.Context, id int64, password string) (repo.User, error) {
	switch {
	case id == 0:
		return repo.User{}, fmt.Errorf("%w: missing 'id'", ErrMissingRequestField)
	case password == "":
		return repo.User{}, fmt.Errorf("%w: missing 'password'", ErrMissingRequestField)
	}

	bp, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return repo.User{}, multierr.Append(ErrInternalServerError, ErrPasswordHashing)
	}
	hashedPassword := string(bp)

	u, err := b.repo.UpdateUserPassword(ctx, repo.UpdateUserPasswordParams{
		ID:             id,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return repo.User{}, fmt.Errorf("%w: no user with 'id': %d", ErrNotFound, id)
		}

		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Code {
			case "23505":
				return repo.User{}, multierr.Append(ErrDuplicatedValue, pqErr)
			}
		}

		return repo.User{}, multierr.Append(
			ErrInternalServerError,
			fmt.Errorf("failed to update user's 'password': %w", err),
		)
	}

	return u, nil
}

func (b *basicUserService) UpdateUserInfo(ctx context.Context, id int64, user repo.User) (repo.User, error) {
	switch {
	case id == 0:
		return repo.User{}, fmt.Errorf("%w: missing 'id'", ErrMissingRequestField)
	}

	u, err := b.repo.UpdateUserInfo(ctx, repo.UpdateUserInfoParams{
		ID:          id,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		BirthDay:    user.BirthDay,
		Gender:      user.Gender,
		PhoneNumber: user.PhoneNumber,
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return repo.User{}, fmt.Errorf("%w: no user with 'id': %d", ErrNotFound, id)
		}

		return repo.User{}, multierr.Append(
			ErrInternalServerError,
			fmt.Errorf("failed to update user: %w", err),
		)
	}

	return u, nil
}

func (b *basicUserService) DeleteUserSoft(ctx context.Context, id int64) error {
	if id == 0 {
		return fmt.Errorf("%w: missing 'id'", ErrMissingRequestField)
	}

	err := b.repo.DeleteUserSoft(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("%w: no user with 'id': %d", ErrNotFound, id)
		}

		return multierr.Append(
			ErrInternalServerError,
			fmt.Errorf("failed to delete user: %w", err),
		)
	}

	return nil
}
func (b *basicUserService) RecoverUser(ctx context.Context, id int64) (repo.User, error) {
	if id == 0 {
		return repo.User{}, fmt.Errorf("%w: missing 'id'", ErrMissingRequestField)
	}

	u, err := b.repo.RecoverUser(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return repo.User{}, fmt.Errorf("%w: no user with 'id': %d", ErrNotFound, id)
		}

		return repo.User{}, multierr.Append(
			ErrInternalServerError,
			fmt.Errorf("failed to recover user: %w", err),
		)
	}

	return u, nil
}

func (b *basicUserService) DeleteUserPermanent(ctx context.Context, id int64) error {
	if id == 0 {
		return fmt.Errorf("%w: missing 'id'", ErrMissingRequestField)
	}

	err := b.repo.DeleteUserPermanent(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("%w: no user with 'id': %d", ErrNotFound, id)
		}

		return multierr.Append(
			ErrInternalServerError,
			fmt.Errorf("failed to delete user: %w", err),
		)
	}

	return nil
}

func (b *basicUserService) VerifyPassword(ctx context.Context, id int64, password string) error {
	switch {
	case id == 0:
		return fmt.Errorf("%w: missing 'id'", ErrMissingRequestField)
	case password == "":
		return fmt.Errorf("%w: missing 'password'", ErrMissingRequestField)
	}

	hashedPassword, err := b.repo.GetHashedPassword(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("%w: no user with 'id': %d", ErrNotFound, id)
		}

		return multierr.Append(
			ErrInternalServerError,
			fmt.Errorf("failed to retrieve user's 'password': %w", err),
		)
	}

	// verify
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return ErrWrongPassword
	}

	return nil
}

// NewBasicUserService returns a naive, stateless implementation of UserService.
func NewBasicUserService(repo repo.Querier) UserService {
	return &basicUserService{repo}
}

// New returns a UserService with all of the expected middleware wired in.
func New(repo repo.Querier, middleware []Middleware) UserService {
	var svc = NewBasicUserService(repo)
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
