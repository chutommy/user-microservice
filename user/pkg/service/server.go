package service

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/chutified/booking-terminal/user/pkg/grpc/userpb"
	"github.com/chutified/booking-terminal/user/pkg/repo"
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
}

// NewUserServer constructs a UserServer.
func NewUserServer(repo repo.Querier) *UserServer {
	return &UserServer{
		repo: repo,
	}
}

func (u *UserServer) RegisterUser(ctx context.Context, req *userpb.RegisterUserRequest) (*userpb.RegisterUserResponse, error) {
	user := req.GetUser()

	switch {
	case user.GetEmail() == "":
		return nil, status.Errorf(codes.InvalidArgument, "%v: 'email' field", ErrEmptyField)
	case user.GetPassword() == "":
		return nil, status.Errorf(codes.InvalidArgument, "%v: 'password' field", ErrEmptyField)
	case user.GetFirstName() == "":
		return nil, status.Errorf(codes.InvalidArgument, "%v: 'first_name' field", ErrEmptyField)
	case user.GetLastName() == "":
		return nil, status.Errorf(codes.InvalidArgument, "%v: 'last_name' field", ErrEmptyField)
	}

	// process password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "fail to hash password")
	}

	// process birthday
	var bdTime sql.NullTime
	if bd := user.GetBirthday(); bd != "" {
		parsedBD, err := time.Parse(ShortForm, bd)
		if err != nil {
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

		return nil, status.Errorf(code, "cannot create a new user: %v", err)
	}

	// construct response
	resp := &userpb.RegisterUserResponse{
		Id: newUser.ID.String(),
	}

	return resp, nil
}

func (u *UserServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	id := req.GetId()
	if id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "%v: 'id' field", ErrEmptyField)
	}

	// parse ID
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id '%v': does not follow UUID pattern", id)
	}

	// retrieve user
	user, err := u.repo.GetUser(ctx, uid)
	if err != nil {
		code := codes.Internal

		if errors.Is(err, sql.ErrNoRows) {
			code = codes.NotFound
		}

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
			Birthday:  user.BirthDay.Time.Format(ShortForm),
		},
	}

	return resp, nil
}

func (u *UserServer) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	panic("implement me")
}

func (u *UserServer) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	id := req.GetId()
	if id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "%v: 'id' field", ErrEmptyField)
	}

	// parse ID
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid id '%v': does not follow UUID pattern", id)
	}

	// remove user
	affected, err := u.repo.DeleteUser(ctx, uid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to retrieve user with id: %s", id)
	}
	if affected != 0 {
		return nil, status.Errorf(codes.NotFound, "failed to retrieve user with id: %s", id)
	}

	// construct response
	resp := &userpb.DeleteUserResponse{
		Id: uid.String(),
	}

	return resp, nil
}
