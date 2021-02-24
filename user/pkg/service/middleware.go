package service

import (
	"context"

	"user/pkg/repo"

	"github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(UserService) UserService

type loggingMiddleware struct {
	logger log.Logger
	next   UserService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a UserService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next UserService) UserService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) AddGender(ctx context.Context, title string) (r0 repo.Gender, err error) {
	defer func() {
		_ = l.logger.Log("method", "AddGender", "title", title, "err", err)
	}()
	return l.next.AddGender(ctx, title)
}

func (l loggingMiddleware) GetGender(ctx context.Context, id int16) (r0 repo.Gender, err error) {
	defer func() {
		_ = l.logger.Log("method", "GetGender", "id", id, "err", err)
	}()
	return l.next.GetGender(ctx, id)
}

func (l loggingMiddleware) ListGenders(ctx context.Context) (r0 []repo.Gender, err error) {
	defer func() {
		_ = l.logger.Log("method", "ListGenders", "err", err)
	}()
	return l.next.ListGenders(ctx)
}

func (l loggingMiddleware) RemoveGender(ctx context.Context, id int16) (err error) {
	defer func() {
		_ = l.logger.Log("method", "RemoveGender", "id", id, "err", err)
	}()
	return l.next.RemoveGender(ctx, id)
}

func (l loggingMiddleware) CreateUser(ctx context.Context, user repo.User) (r0 repo.User, err error) {
	defer func() {
		_ = l.logger.Log("method", "CreateUser", "user", user, "err", err)
	}()
	return l.next.CreateUser(ctx, user)
}

func (l loggingMiddleware) GetUserByID(ctx context.Context, id int64) (r0 repo.User, err error) {
	defer func() {
		_ = l.logger.Log("method", "GetUserByID", "id", id, "err", err)
	}()
	return l.next.GetUserByID(ctx, id)
}

func (l loggingMiddleware) GetUserByEmail(ctx context.Context, email string) (r0 repo.User, err error) {
	defer func() {
		_ = l.logger.Log("method", "GetUserByEmail", "email", email, "err", err)
	}()
	return l.next.GetUserByEmail(ctx, email)
}

func (l loggingMiddleware) UpdateUserEmail(ctx context.Context, id int64, email string) (r0 repo.User, err error) {
	defer func() {
		_ = l.logger.Log("method", "UpdateUserEmail", "id", id, "email", email, "err", err)
	}()
	return l.next.UpdateUserEmail(ctx, id, email)
}

func (l loggingMiddleware) UpdateUserPassword(ctx context.Context, id int64, password string) (r0 repo.User, err error) {
	defer func() {
		_ = l.logger.Log("method", "UpdateUserPassword", "id", id, "password", password, "err", err)
	}()
	return l.next.UpdateUserPassword(ctx, id, password)
}

func (l loggingMiddleware) UpdateUserInfo(ctx context.Context, id int64, user repo.User) (r0 repo.User, err error) {
	defer func() {
		_ = l.logger.Log("method", "UpdateUserInfo", "id", id, "user", user, "err", err)
	}()
	return l.next.UpdateUserInfo(ctx, id, user)
}

func (l loggingMiddleware) DeleteUserSoft(ctx context.Context, id int64) (err error) {
	defer func() {
		_ = l.logger.Log("method", "DeleteUserSoft", "id", id, "err", err)
	}()
	return l.next.DeleteUserSoft(ctx, id)
}

func (l loggingMiddleware) RecoverUser(ctx context.Context, id int64) (r0 repo.User, err error) {
	defer func() {
		_ = l.logger.Log("method", "RecoverUser", "id", id, "err", err)
	}()
	return l.next.RecoverUser(ctx, id)
}

func (l loggingMiddleware) DeleteUserPermanent(ctx context.Context, id int64) (err error) {
	defer func() {
		_ = l.logger.Log("method", "DeleteUserPermanent", "id", id, "err", err)
	}()
	return l.next.DeleteUserPermanent(ctx, id)
}

func (l loggingMiddleware) VerifyPassword(ctx context.Context, id int64, password string) (err error) {
	defer func() {
		_ = l.logger.Log("method", "VerifyPassword", "id", id, "password", password, "err", err)
	}()
	return l.next.VerifyPassword(ctx, id, password)
}
