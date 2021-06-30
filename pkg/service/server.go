package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/lib/pq"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/chutommy/user-microservice/pkg/grpc/userpb"
	"github.com/chutommy/user-microservice/pkg/repo"
)

var (
	// ShortForm is a format of a birthday date.
	ShortForm = "2006-Jan-02"

	// ErrMissingArgument is returned if the argument is invalid because of an empty
	// mandatory field.
	ErrEmptyField = errors.New("required field has empty value")
)

// UserServer implements userpb.UserServiceServer.
type UserServer struct {
	userpb.UnimplementedUserServiceServer

	repo repo.Querier

	// TODO: add logger middleware
}

// NewUserServer constructs a UserServer.
func NewUserServer(repo repo.Querier) *UserServer {
	return &UserServer{
		repo: repo,
	}
}

func (u *UserServer) RegisterUser(ctx context.Context, req *userpb.RegisterUserRequest) (*userpb.RegisterUserResponse, error) {
	logger := ctxzap.Extract(ctx)
	user := req.GetUser()

	switch {
	case user.GetEmail() == "":
		logger.Info("empty email")
		return nil, status.Errorf(codes.InvalidArgument, "%v: 'email' field", ErrEmptyField)
	case user.GetPassword() == "":
		logger.Info("empty password")
		return nil, status.Errorf(codes.InvalidArgument, "%v: 'password' field", ErrEmptyField)
	case user.GetFirstName() == "":
		logger.Info("empty first name")
		return nil, status.Errorf(codes.InvalidArgument, "%v: 'first_name' field", ErrEmptyField)
	case user.GetLastName() == "":
		logger.Info("empty last name")
		return nil, status.Errorf(codes.InvalidArgument, "%v: 'last_name' field", ErrEmptyField)
	}

	// process password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("failed to hash the password", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "fail to hash password")
	}

	// process birthday
	var bdTime sql.NullTime
	if bd := user.GetBirthday(); bd != "" {
		parsedBD, err := time.Parse(ShortForm, bd)
		if err != nil {
			logger.Info("failed to parse birthday", zap.Error(err))
			return nil, status.Errorf(codes.InvalidArgument, "field time is in unsupported format: %v instead of %v", err, ShortForm)
		}

		bdTime = sql.NullTime{
			Time:  parsedBD,
			Valid: true,
		}
	}

	// build argument
	arg := repo.CreateUserParams{
		ID:    uuid.New(),
		Email: user.GetEmail(),
		PhoneNumber: sql.NullString{
			String: user.GetPhone(),
			Valid:  user.GetPhone() != "",
		},
		HashedPassword: string(hashedPassword),
		FirstName:      user.GetFirstName(),
		LastName:       user.GetLastName(),
		Gender:         int16(user.GetGender()),
		BirthDay:       bdTime,
	}

	// store user
	newUser, err := u.repo.CreateUser(ctx, arg)
	if err != nil {
		code := codes.Internal

		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				code = codes.AlreadyExists
			}
		}

		logger.Error("failed to create a new user", zap.Error(err))
		return nil, status.Errorf(code, "cannot create a new user: %v", err)
	}

	// construct response
	resp := &userpb.RegisterUserResponse{
		Id: newUser.ID.String(),
	}

	return resp, nil
}

func (u *UserServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	logger := ctxzap.Extract(ctx)

	id := req.GetId()
	if id == "" {
		logger.Info("empty id")
		return nil, status.Errorf(codes.InvalidArgument, "%v: 'id' field", ErrEmptyField)
	}

	// parse ID
	uid, err := uuid.Parse(id)
	if err != nil {
		logger.Info("invalid uuid", zap.String("uuid", id), zap.Error(err))
		return nil, status.Errorf(codes.InvalidArgument, "invalid id '%v': does not follow UUID pattern", id)
	}

	// retrieve user
	user, err := u.repo.GetUser(ctx, uid)
	if err != nil {
		code := codes.Internal

		if errors.Is(err, sql.ErrNoRows) {
			code = codes.NotFound
		}

		logger.Error("retrieve user", zap.Error(err))
		return nil, status.Errorf(code, "failed to retrieve user with id: %s", id)
	}

	// construct response
	resp := &userpb.GetUserResponse{
		User: &userpb.User{
			Id:        user.ID.String(),
			Email:     user.Email,
			Phone:     user.PhoneNumber.String,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Gender:    userpb.User_Gender(user.Gender),
		},
	}
	if user.BirthDay.Valid {
		resp.User.Birthday = user.BirthDay.Time.Format(ShortForm)
	}

	return resp, nil
}

func (u *UserServer) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	logger := ctxzap.Extract(ctx)
	user := req.GetUser()

	id := req.GetId()
	if id == "" {
		logger.Info("empty id")
		return nil, status.Errorf(codes.InvalidArgument, "%v: 'id' field", ErrEmptyField)
	}

	// process id
	uid, err := uuid.Parse(id)
	if err != nil {
		logger.Info("invalid uuid id", zap.Error(err))
		return nil, status.Errorf(codes.InvalidArgument, "invalid id '%v': does not follow UUID pattern", id)
	}

	// process password
	var hashedPassword []byte
	if p := user.GetPassword(); p != "" {
		hashedPassword, err = bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
		if err != nil {
			logger.Error("invalid hashed password", zap.Error(err))
			return nil, status.Errorf(codes.Internal, "fail to hash password")
		}
	}

	// process birthday
	var bdTime time.Time
	if bd := user.GetBirthday(); bd != "" {
		bdTime, err = time.Parse(ShortForm, bd)
		if err != nil {
			logger.Info("invalid birthday format")
			return nil, status.Errorf(codes.InvalidArgument, "field time is in unsupported format: %v instead of %v", err, ShortForm)
		}
	}

	// construct na argument
	arg := repo.UpdateUserParams{
		ID: uid,

		Email:          user.GetEmail(),
		PhoneNumber:    user.GetPhone(),
		HashedPassword: string(hashedPassword),
		FirstName:      user.GetFirstName(),
		LastName:       user.GetLastName(),
		Gender:         int16(user.GetGender()),
		BirthDay:       bdTime,
	}

	// update user
	updUser, err := u.repo.UpdateUser(ctx, arg)
	if err != nil {
		code := codes.Internal

		if errors.Is(err, sql.ErrNoRows) {
			code = codes.NotFound
		}

		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				code = codes.AlreadyExists
			}
		}

		logger.Error("failed to update user", zap.Error(err))
		return nil, status.Errorf(code, "failed to update user with an id '%s'", id)
	}

	// construct a response
	resp := &userpb.UpdateUserResponse{
		Id: updUser.ID.String(),
	}

	return resp, nil
}

func (u *UserServer) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	logger := ctxzap.Extract(ctx)

	id := req.GetId()
	if id == "" {
		logger.Info("empty id")
		return nil, status.Errorf(codes.InvalidArgument, "%v: 'id' field", ErrEmptyField)
	}

	// parse ID
	uid, err := uuid.Parse(id)
	if err != nil {
		logger.Info("infalid uuid id", zap.Error(err))
		return nil, status.Errorf(codes.InvalidArgument, "invalid id '%v': does not follow UUID pattern", id)
	}

	// remove user
	affected, err := u.repo.DeleteUser(ctx, uid)
	if err != nil || affected != 1 {
		code := codes.Internal

		if affected == 0 && err == nil {
			code = codes.NotFound
		}

		logger.Info("failed to delete user", zap.Error(err))
		return nil, status.Errorf(code, "failed to delete user with id: %s", id)
	}

	// construct response
	resp := &userpb.DeleteUserResponse{
		Id: uid.String(),
	}

	return resp, nil
}
