package service_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/chutified/booking-terminal/user/pkg/grpc/userpb"
	"github.com/chutified/booking-terminal/user/pkg/mocks"
	"github.com/chutified/booking-terminal/user/pkg/repo"
	"github.com/chutified/booking-terminal/user/pkg/service"
	"github.com/chutified/booking-terminal/user/pkg/util"
)

func randomUser() *userpb.User {
	// construct a random user
	return &userpb.User{
		Id:        uuid.New().String(),
		Email:     util.RandomEmail(),
		Phone:     util.RandomPhoneNumber(),
		Password:  util.RandomPassword(),
		FirstName: util.RandomName(),
		LastName:  util.RandomName(),
		Gender:    userpb.User_Gender(util.RandomInt(0, 3)),
		Birthday:  time.Now().Add(-200000 * time.Hour).Format(service.ShortForm),
	}
}

func TestUserServer_RegisterUser(t *testing.T) {
	t.Parallel()

	u1 := randomUser()
	u1p, _ := bcrypt.GenerateFromPassword([]byte(u1.Password), bcrypt.DefaultCost)
	u1t, err := time.Parse(service.ShortForm, u1.Birthday)
	require.NoError(t, err)

	tests := []struct {
		name      string
		buildRepo func(q *mocks.Querier)
		argUser   *userpb.User
		expID     string
		expCode   codes.Code
	}{
		{
			name: "ok",
			buildRepo: func(q *mocks.Querier) {
				q.On(
					"CreateUser",
					mock.Anything,
					mock.AnythingOfType("repo.CreateUserParams"),
				).Return(repo.User{
					ID:    uuid.MustParse(u1.Id),
					Email: u1.Email,
					PhoneNumber: sql.NullString{
						String: u1.Phone,
						Valid:  true,
					},
					HashedPassword: string(u1p),
					FirstName:      u1.FirstName,
					LastName:       u1.LastName,
					Gender:         int16(u1.Gender),
					BirthDay: sql.NullTime{
						Time:  u1t,
						Valid: true,
					},
					UpdatedAt: sql.NullTime{},
					CreatedAt: time.Now(),
				}, nil).Once()
			},
			argUser: u1,
			expID:   u1.Id,
			expCode: codes.OK,
		},
		{
			name: "empty optional fields",
			buildRepo: func(q *mocks.Querier) {
				q.On(
					"CreateUser",
					mock.Anything,
					mock.AnythingOfType("repo.CreateUserParams"),
				).Return(repo.User{
					ID:             uuid.MustParse(u1.Id),
					Email:          u1.Email,
					PhoneNumber:    sql.NullString{},
					HashedPassword: string(u1p),
					FirstName:      u1.FirstName,
					LastName:       u1.LastName,
					Gender:         0,
					BirthDay:       sql.NullTime{},
					UpdatedAt:      sql.NullTime{},
					CreatedAt:      time.Now(),
				}, nil).Once()
			},
			argUser: &userpb.User{
				Email:     u1.Email,
				Phone:     "",
				Password:  u1.Password,
				FirstName: u1.FirstName,
				LastName:  u1.LastName,
				Gender:    userpb.User_Gender(0),
				Birthday:  "",
			},
			expID:   u1.Id,
			expCode: codes.OK,
		},
		{
			name:      "empty email",
			buildRepo: func(q *mocks.Querier) {},
			argUser: &userpb.User{
				Email:     "",
				Phone:     u1.Phone,
				Password:  u1.Password,
				FirstName: u1.FirstName,
				LastName:  u1.LastName,
				Gender:    u1.Gender,
				Birthday:  u1.Birthday,
			},
			expID:   "",
			expCode: codes.InvalidArgument,
		},
		{
			name:      "empty password",
			buildRepo: func(q *mocks.Querier) {},
			argUser: &userpb.User{
				Email:     u1.Email,
				Phone:     u1.Phone,
				Password:  "",
				FirstName: u1.FirstName,
				LastName:  u1.LastName,
				Gender:    u1.Gender,
				Birthday:  u1.Birthday,
			},
			expID:   "",
			expCode: codes.InvalidArgument,
		},
		{
			name:      "empty first name",
			buildRepo: func(q *mocks.Querier) {},
			argUser: &userpb.User{
				Email:     u1.Email,
				Phone:     u1.Phone,
				Password:  u1.Password,
				FirstName: "",
				LastName:  u1.LastName,
				Gender:    u1.Gender,
				Birthday:  u1.Birthday,
			},
			expID:   "",
			expCode: codes.InvalidArgument,
		},
		{
			name:      "empty last name",
			buildRepo: func(q *mocks.Querier) {},
			argUser: &userpb.User{
				Email:     u1.Email,
				Phone:     u1.Phone,
				Password:  u1.Password,
				FirstName: u1.FirstName,
				LastName:  "",
				Gender:    u1.Gender,
				Birthday:  u1.Birthday,
			},
			expID:   "",
			expCode: codes.InvalidArgument,
		},
		{
			name: "unique key violation",
			buildRepo: func(q *mocks.Querier) {
				q.On(
					"CreateUser",
					mock.Anything,
					mock.AnythingOfType("repo.CreateUserParams"),
				).Return(repo.User{}, &pq.Error{Code: "23505"}).Once()
			},
			argUser: u1,
			expID:   "",
			expCode: codes.AlreadyExists,
		},
		{
			name: "connection error",
			buildRepo: func(q *mocks.Querier) {
				q.On(
					"CreateUser",
					mock.Anything,
					mock.AnythingOfType("repo.CreateUserParams"),
				).Return(repo.User{}, sql.ErrConnDone).Once()
			},
			argUser: u1,
			expID:   "",
			expCode: codes.Internal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// construct server
			mockRepo := new(mocks.Querier)
			tt.buildRepo(mockRepo)
			server := service.NewUserServer(mockRepo)

			// build request
			req := &userpb.RegisterUserRequest{User: tt.argUser}

			// test method
			resp, err := server.RegisterUser(context.Background(), req)
			if tt.expCode == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, resp)

				require.NotEmpty(t, resp.Id)
				require.Equal(t, tt.expID, resp.Id)
			} else {
				require.Error(t, err)
				require.Nil(t, resp)

				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tt.expCode, st.Code())
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestUserServer_GetUser(t *testing.T) {
	t.Parallel()

	// build a random user for testing purpose
	u1 := randomUser()
	u1p, _ := bcrypt.GenerateFromPassword([]byte(u1.Password), bcrypt.DefaultCost)
	u1t, err := time.Parse(service.ShortForm, u1.Birthday)
	require.NoError(t, err)

	tests := []struct {
		name      string
		buildRepo func(q *mocks.Querier)
		inpID     string
		expUser   *userpb.User
		expCode   codes.Code
	}{
		{
			name: "ok",
			buildRepo: func(q *mocks.Querier) {
				q.On("GetUser", mock.Anything, mock.Anything).Return(repo.User{
					ID:    uuid.MustParse(u1.Id),
					Email: u1.Email,
					PhoneNumber: sql.NullString{
						String: u1.Phone,
						Valid:  true,
					},
					HashedPassword: string(u1p),
					FirstName:      u1.FirstName,
					LastName:       u1.LastName,
					Gender:         int16(u1.Gender),
					BirthDay: sql.NullTime{
						Time:  u1t,
						Valid: true,
					},
					UpdatedAt: sql.NullTime{},
					CreatedAt: time.Now().Add(-1000 * time.Hour),
				}, nil)
			},
			inpID:   u1.Id,
			expUser: u1,
			expCode: codes.OK,
		},
		{
			name:      "empty id",
			buildRepo: func(q *mocks.Querier) {},
			inpID:     "",
			expUser:   nil,
			expCode:   codes.InvalidArgument,
		},
		{
			name:      "invalid id",
			buildRepo: func(q *mocks.Querier) {},
			inpID:     "invalid_uuid",
			expUser:   nil,
			expCode:   codes.InvalidArgument,
		},
		{
			name: "user not found",
			buildRepo: func(q *mocks.Querier) {
				q.On("GetUser", mock.Anything, mock.Anything).Return(repo.User{}, sql.ErrNoRows)
			},
			inpID:   u1.Id,
			expUser: &userpb.User{},
			expCode: codes.NotFound,
		},
		{
			name: "internal error",
			buildRepo: func(q *mocks.Querier) {
				q.On("GetUser", mock.Anything, mock.Anything).Return(repo.User{}, sql.ErrConnDone)
			},
			inpID:   u1.Id,
			expUser: &userpb.User{},
			expCode: codes.Internal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// construct a mock server
			mockRepo := new(mocks.Querier)
			tt.buildRepo(mockRepo)
			server := service.NewUserServer(mockRepo)

			arg := &userpb.GetUserRequest{Id: tt.inpID}

			resp, err := server.GetUser(context.Background(), arg)
			if tt.expCode == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, resp.User)

				tt.expUser.Password = ""
				require.Equal(t, tt.expUser, resp.User)
			} else {
				require.Error(t, err)
				require.Nil(t, resp)

				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tt.expCode, st.Code())
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestUserServer_DeleteUser(t *testing.T) {
	t.Parallel()

	// build a random user for testing purpose
	u1 := randomUser()

	tests := []struct {
		name      string
		buildRepo func(q *mocks.Querier)
		inpID     string
		expID     string
		expCode   codes.Code
	}{
		{
			name: "ok",
			buildRepo: func(q *mocks.Querier) {
				q.On("DeleteUser", mock.Anything, uuid.MustParse(u1.Id)).Return(int64(1), nil)
			},
			inpID:   u1.Id,
			expID:   u1.Id,
			expCode: codes.OK,
		},
		{
			name:      "missing id",
			buildRepo: func(q *mocks.Querier) {},
			inpID:     "",
			expID:     "",
			expCode:   codes.InvalidArgument,
		},
		{
			name:      "invalid id",
			buildRepo: func(q *mocks.Querier) {},
			inpID:     "invalid",
			expID:     "",
			expCode:   codes.InvalidArgument,
		},
		{
			name: "not found",
			buildRepo: func(q *mocks.Querier) {
				q.On("DeleteUser", mock.Anything, uuid.MustParse(u1.Id)).Return(int64(0), nil)
			},
			inpID:   u1.Id,
			expID:   "",
			expCode: codes.NotFound,
		},
		{
			name: "affected more id",
			buildRepo: func(q *mocks.Querier) {
				q.On("DeleteUser", mock.Anything, uuid.MustParse(u1.Id)).Return(int64(2), nil)
			},
			inpID:   u1.Id,
			expID:   "",
			expCode: codes.Internal,
		},
		{
			name: "internal error",
			buildRepo: func(q *mocks.Querier) {
				q.On("DeleteUser", mock.Anything, uuid.MustParse(u1.Id)).Return(int64(0), sql.ErrConnDone)
			},
			inpID:   u1.Id,
			expID:   "",
			expCode: codes.Internal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// construct a mock server
			mockRepo := new(mocks.Querier)
			tt.buildRepo(mockRepo)
			server := service.NewUserServer(mockRepo)

			arg := &userpb.DeleteUserRequest{Id: tt.inpID}

			resp, err := server.DeleteUser(context.Background(), arg)
			if tt.expCode == codes.OK {
				require.NoError(t, err)
				require.NotNil(t, resp)
				require.Equal(t, tt.expID, resp.Id)
			} else {
				require.Error(t, err)
				require.Nil(t, resp)

				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tt.expCode, st.Code())
				require.Empty(t, tt.expID)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
