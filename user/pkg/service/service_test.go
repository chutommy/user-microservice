package service_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/lib/pq"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"

	"user/pkg/mocks"
	"user/pkg/repo"
	"user/pkg/service"
	"user/pkg/util"
)

func TestBasicUserService_AddGender(t *testing.T) {
	g := repo.Gender{
		ID:    int16(util.RandomInt(0, 1024)),
		Title: util.RandomShortString(),
	}

	tests := []struct {
		name      string
		title     string
		buildStub func(*mocks.Querier)
		expGender repo.Gender
		expErr    error
	}{
		{
			name:  "OK",
			title: g.Title,
			buildStub: func(q *mocks.Querier) {
				q.On("CreateGender", mock.Anything, g.Title).
					Return(g, nil)
			},
			expGender: g,
			expErr:    nil,
		},
		{
			name:      "EmptyTitle",
			title:     "",
			buildStub: func(q *mocks.Querier) {},
			expGender: repo.Gender{},
			expErr:    service.ErrEmptyTitleField,
		},
		{
			name:  "UniqueViolation",
			title: g.Title,
			buildStub: func(q *mocks.Querier) {
				q.On("CreateGender", mock.Anything, g.Title).
					Return(repo.Gender{}, pq.Error{
						Code:    "23505",
						Message: "unique_violation",
					})
			},
			expGender: repo.Gender{},
			expErr:    service.ErrUniqueGenderViolation,
		},
		{
			name:  "InternalError",
			title: g.Title,
			buildStub: func(q *mocks.Querier) {
				q.On("CreateGender", mock.Anything, g.Title).
					Return(repo.Gender{}, sql.ErrConnDone)
			},
			expGender: repo.Gender{},
			expErr:    service.ErrInternalDBError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// build service
			mockRepo := new(mocks.Querier)
			tt.buildStub(mockRepo)
			svc := service.NewBasicUserService(mockRepo)

			// serve
			gOut, err := svc.AddGender(context.Background(), tt.title)
			require.Equal(t, tt.expGender, gOut)
			require.True(t, errors.Is(err, tt.expErr))

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestBasicUserService_GetGender(t *testing.T) {
	g := repo.Gender{
		ID:    int16(util.RandomInt(0, 1024)),
		Title: util.RandomShortString(),
	}

	tests := []struct {
		name      string
		buildStub func(*mocks.Querier)
		inpID     int16
		expGender repo.Gender
		expErr    error
	}{
		{
			name: "OK",
			buildStub: func(q *mocks.Querier) {
				q.On("GetGender", mock.Anything, g.ID).Return(g, nil)
			},
			inpID:     g.ID,
			expGender: g,
			expErr:    nil,
		},
		{
			name:      "InvalidRequest",
			buildStub: func(q *mocks.Querier) {},
			inpID:     0,
			expGender: repo.Gender{},
			expErr:    service.ErrEmptySearchKeys,
		},
		{
			name: "NotFound",
			buildStub: func(q *mocks.Querier) {
				q.On("GetGender", mock.Anything, g.ID).Return(repo.Gender{}, sql.ErrNoRows)
			},
			inpID:     g.ID,
			expGender: repo.Gender{},
			expErr:    service.ErrNotFound,
		},
		{
			name: "InternalError",
			buildStub: func(q *mocks.Querier) {
				q.On("GetGender", mock.Anything, g.ID).Return(repo.Gender{}, sql.ErrConnDone)
			},
			inpID:     g.ID,
			expGender: repo.Gender{},
			expErr:    service.ErrInternalDBError,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// build service
			mockRepo := new(mocks.Querier)
			test.buildStub(mockRepo)
			svc := service.NewBasicUserService(mockRepo)

			// serve
			g1, err := svc.GetGender(context.Background(), test.inpID)
			require.Equal(t, test.expGender, g1)
			require.True(t, errors.Is(err, test.expErr))

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestBasicUserService_ListGenders(t *testing.T) {
	genders := []repo.Gender{
		{
			ID:    int16(util.RandomInt(0, 1024)),
			Title: util.RandomShortString(),
		},
		{
			ID:    int16(util.RandomInt(1025, 2048)),
			Title: util.RandomString(),
		},
	}

	tests := []struct {
		name       string
		buildStub  func(q *mocks.Querier)
		expGenders []repo.Gender
		expErr     error
	}{
		{
			name: "OK",
			buildStub: func(q *mocks.Querier) {
				q.On("ListGenders", mock.Anything).
					Return(genders, nil)
			},
			expGenders: genders,
			expErr:     nil,
		},
		{
			name: "InternalError",
			buildStub: func(q *mocks.Querier) {
				q.On("ListGenders", mock.Anything).
					Return([]repo.Gender{}, sql.ErrConnDone)
			},
			expGenders: []repo.Gender{},
			expErr:     service.ErrInternalDBError,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// build service
			mockRepo := new(mocks.Querier)
			test.buildStub(mockRepo)
			svc := service.NewBasicUserService(mockRepo)

			// serve
			genders1, err := svc.ListGenders(context.Background())
			require.Equal(t, test.expGenders, genders1)
			require.True(t, errors.Is(err, test.expErr))

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestBasicUserService_RemoveGender(t *testing.T) {
	g := repo.Gender{
		ID:    int16(util.RandomInt(0, 1024)),
		Title: util.RandomShortString(),
	}

	tests := []struct {
		name      string
		buildStub func(q *mocks.Querier)
		inpID     int16
		expErr    error
	}{
		{
			name: "OK",
			buildStub: func(q *mocks.Querier) {
				q.On("DeleteGender", mock.Anything, g.ID).Return(nil)
			},
			inpID:  g.ID,
			expErr: nil,
		},
		{
			name: "NotFound",
			buildStub: func(q *mocks.Querier) {
				q.On("DeleteGender", mock.Anything, g.ID).Return(sql.ErrNoRows)
			},
			inpID:  g.ID,
			expErr: service.ErrNotFound,
		},
		{
			name:      "EmptyID",
			buildStub: func(q *mocks.Querier) {},
			inpID:     0,
			expErr:    service.ErrEmptyID,
		},
		{
			name: "InternalError",
			buildStub: func(q *mocks.Querier) {
				q.On("DeleteGender", mock.Anything, g.ID).Return(sql.ErrConnDone)
			},
			inpID:  g.ID,
			expErr: service.ErrInternalDBError,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// build service
			mockRepo := new(mocks.Querier)
			test.buildStub(mockRepo)
			svc := service.NewBasicUserService(mockRepo)

			// serve
			err := svc.RemoveGender(context.Background(), test.inpID)
			require.True(t, errors.Is(err, test.expErr))

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestBasicUserService_CreateUser(t *testing.T) {
	// get date
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, &time.Location{})

	// prepare password
	password := util.RandomString()
	hashedPasswordB, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashedPassword := string(hashedPasswordB)
	require.NoError(t, err)

	user := repo.User{
		ID:             util.RandomInt(0, 1024),
		HashedPassword: hashedPassword,
		FirstName:      util.RandomShortString(),
		LastName:       util.RandomLongString(),
		BirthDay: sql.NullTime{
			Time:  today.AddDate(-20, 0, 0),
			Valid: true,
		},
		Gender: int16(util.RandomInt(0, 256)),
		Email:  util.RandomEmail(),
		PhoneNumber: sql.NullString{
			String: util.RandomPhoneNumber(),
			Valid:  true,
		},
		UpdatedAt: sql.NullTime{Valid: false},
		DeletedAt: sql.NullTime{Valid: false},
		CreatedAt: now,
	}

	type inp struct {
		password    string
		firstName   string
		lastName    string
		gender      int16
		email       string
		phoneNumber string
		birthday    time.Time
	}

	type exp struct {
		user repo.User
		err  error
	}

	tests := []struct {
		name      string
		buildStub func(q *mocks.Querier)
		inp       inp
		exp       exp
	}{
		{
			name: "OK",
			buildStub: func(q *mocks.Querier) {
				q.On("CreateUser", mock.Anything, repo.CreateUserParams{
					HashedPassword: user.HashedPassword,
					FirstName:      user.FirstName,
					LastName:       user.LastName,
					BirthDay:       user.BirthDay,
					Gender:         user.Gender,
					Email:          user.Email,
					PhoneNumber:    user.PhoneNumber,
				}, nil).Return(user, nil)
			},
			inp: inp{
				password:    password,
				firstName:   user.FirstName,
				lastName:    user.LastName,
				gender:      user.Gender,
				email:       user.Email,
				phoneNumber: user.PhoneNumber.String,
				birthday:    user.BirthDay.Time,
			},
			exp: exp{
				user: user,
				err:  nil,
			},
		},
		{
			name:      "EmptyPassword",
			buildStub: func(q *mocks.Querier) {},
			inp: inp{
				password:    "",
				firstName:   user.FirstName,
				lastName:    user.LastName,
				gender:      user.Gender,
				email:       user.Email,
				phoneNumber: user.PhoneNumber.String,
				birthday:    user.BirthDay.Time,
			},
			exp: exp{
				user: repo.User{},
				err:  service.ErrEmptyPassword,
			},
		},
		{
			name:      "EmptyFirstName",
			buildStub: func(q *mocks.Querier) {},
			inp: inp{
				password:    password,
				firstName:   user.FirstName,
				lastName:    "",
				gender:      user.Gender,
				email:       user.Email,
				phoneNumber: user.PhoneNumber.String,
				birthday:    user.BirthDay.Time,
			},
			exp: exp{
				user: repo.User{},
				err:  service.ErrEmptyFirstName,
			},
		},
		{
			name:      "EmptyLastName",
			buildStub: func(q *mocks.Querier) {},
			inp: inp{
				password:    password,
				firstName:   user.FirstName,
				lastName:    "",
				gender:      user.Gender,
				email:       user.Email,
				phoneNumber: user.PhoneNumber.String,
				birthday:    user.BirthDay.Time,
			},
			exp: exp{
				user: repo.User{},
				err:  service.ErrEmptyLastName,
			},
		},
		{
			name:      "EmptyEmail",
			buildStub: func(q *mocks.Querier) {},
			inp: inp{
				password:    password,
				firstName:   user.FirstName,
				lastName:    user.LastName,
				gender:      user.Gender,
				email:       "",
				phoneNumber: user.PhoneNumber.String,
				birthday:    user.BirthDay.Time,
			},
			exp: exp{
				user: repo.User{},
				err:  service.ErrEmptyEmail,
			},
		},
		{
			name: "DuplicatedEmail",
			buildStub: func(q *mocks.Querier) {
				q.On("CreateUser", mock.Anything, repo.CreateUserParams{
					HashedPassword: user.HashedPassword,
					FirstName:      user.FirstName,
					LastName:       user.LastName,
					BirthDay:       user.BirthDay,
					Gender:         user.Gender,
					Email:          user.Email,
					PhoneNumber:    user.PhoneNumber,
				}, nil).Return(repo.User{}, pq.Error{
					Code:    "23505",
					Message: "duplicate key value violates unique constraint \"users_email_key\"",
				})
			},
			inp: inp{
				password:    password,
				firstName:   user.FirstName,
				lastName:    user.LastName,
				gender:      user.Gender,
				email:       user.Email,
				phoneNumber: user.PhoneNumber.String,
				birthday:    user.BirthDay.Time,
			},
			exp: exp{
				user: repo.User{},
				err:  service.ErrUniqueEmailViolation,
			},
		},
		{
			name: "InternalError",
			buildStub: func(q *mocks.Querier) {
				q.On("CreateUser", mock.Anything, repo.CreateUserParams{
					HashedPassword: user.HashedPassword,
					FirstName:      user.FirstName,
					LastName:       user.LastName,
					BirthDay:       user.BirthDay,
					Gender:         user.Gender,
					Email:          user.Email,
					PhoneNumber:    user.PhoneNumber,
				}, nil).Return(repo.User{}, sql.ErrConnDone)
			},
			inp: inp{
				password:    password,
				firstName:   user.FirstName,
				lastName:    user.LastName,
				gender:      user.Gender,
				email:       user.Email,
				phoneNumber: user.PhoneNumber.String,
				birthday:    user.BirthDay.Time,
			},
			exp: exp{
				user: repo.User{},
				err:  service.ErrInternalDBError,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// build service
			mockRepo := new(mocks.Querier)
			test.buildStub(mockRepo)
			svc := service.NewBasicUserService(mockRepo)

			// serve
			u, err := svc.CreateUser(context.Background(), test.inp.password, test.inp.firstName, test.inp.lastName, test.inp.gender, test.inp.email, test.inp.phoneNumber, test.inp.birthday)
			require.True(t, errors.Is(err, test.exp.err))
			require.Equal(t, test.exp.user, u)

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestBasicUserService_GetUserByID(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: test cases
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}
func TestBasicUserService_GetUserByEmail(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: test cases
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}

func TestBasicUserService_UpdateUserEmail(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: test cases
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}
func TestBasicUserService_UpdateUserPhoneNumber(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: test cases
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}
func TestBasicUserService_UpdateUserPassword(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: test cases
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}

func TestBasicUserService_UpdateUserInfo(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: test cases
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}

func TestBasicUserService_DeleteUserSoft(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: test cases
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}

func TestBasicUserService_RecoverDeletedUser(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: test cases
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}

func TestBasicUserService_DeleteUserPermanent(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: test cases
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}

func TestBasicUserService_VerifyPassword(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: test cases
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

		})
	}
}
