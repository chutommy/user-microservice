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
		Birthday: &userpb.User_Date{
			Year:  int32(util.RandomInt(1900, 2200)),
			Month: int32(util.RandomInt(1, 12)),
			Day:   int32(util.RandomInt(1, 31)),
		},
	}
}

func TestUserServer_RegisterUser(t *testing.T) {
	t.Parallel()

	u1 := randomUser()
	u1p, _ := bcrypt.GenerateFromPassword([]byte(u1.Password), bcrypt.DefaultCost)
	u1t := time.Date(
		int(u1.Birthday.Year),
		time.Month(u1.Birthday.Month),
		int(u1.Birthday.Day),
		0, 0, 0, 0, time.UTC,
	)

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
						String: "",
						Valid:  false,
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
				Birthday:  &userpb.User_Date{},
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
