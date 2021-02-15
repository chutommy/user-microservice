package service_test

import (
	"context"
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
	"user/pkg/mocks"
	"user/pkg/repo"
	"user/pkg/service"
	"user/pkg/util"
)

func TestBasicUserService_AddGender(t *testing.T) {
	g := service.Gender{
		ID:    util.RandomInt(0, 1024),
		Title: util.RandomShortString(),
	}

	tests := []struct {
		name      string
		title     string
		buildStub func(*mocks.Querier)
		expGender service.Gender
		expErr    error
	}{
		{
			name:  "OK",
			title: g.Title,
			buildStub: func(q *mocks.Querier) {
				q.On("CreateGender", mock.Anything, g.Title).
					Return(repo.Gender{
						ID:    int16(g.ID),
						Title: g.Title,
					}, nil)
			},
			expGender: g,
			expErr:    nil,
		},
		{
			name:      "EmptyTitle",
			title:     "",
			buildStub: func(q *mocks.Querier) {},
			expGender: service.Gender{},
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
			expGender: service.Gender{},
			expErr:    service.ErrUniqueGenderViolation,
		},
		{
			name:  "InternalError",
			title: g.Title,
			buildStub: func(q *mocks.Querier) {
				q.On("CreateGender", mock.Anything, g.Title).
					Return(repo.Gender{}, sql.ErrConnDone)
			},
			expGender: service.Gender{},
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
		})
	}
}
