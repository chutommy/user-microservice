package service_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/lib/pq"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

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
func TestBasicUserService_GetUserByUsername(t *testing.T) {
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
