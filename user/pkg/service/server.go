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
	"github.com/chutified/booking-terminal/user/pkg/util"
)

var (
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
	bd := user.GetBirthday()
	bdTime := sql.NullTime{}
	if util.ValidateDate(bd.GetYear(), bd.GetMonth(), bd.GetDay()) {
		bdTime.Time = time.Date(int(bd.GetYear()), time.Month(bd.GetMonth()), int(bd.GetDay()), 0, 0, 0, 0, time.UTC)
		bdTime.Valid = true
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
	panic("implement me")
}

func (u *UserServer) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	panic("implement me")
}

func (u *UserServer) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	panic("implement me")
}
