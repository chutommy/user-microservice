package repo_test

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
	"user/pkg/repo"
	"user/pkg/util"
)

func createRandomGender(t *testing.T) repo.Gender {
	t.Helper()

	// create gender
	title := util.RandomShortString()
	g1, err := testQueries.CreateGender(context.Background(), title)
	require.NoError(t, err)
	require.NotEmpty(t, g1)

	// check db
	param := repo.GetGenderParams{ID: g1.ID}
	g2, err := testQueries.GetGender(context.Background(), param)
	require.NoError(t, err)
	require.NotEmpty(t, g2)

	// check g1 == g2
	require.Equal(t, g1.Title, title)
	require.Equal(t, g1, g2)

	return g1
}

func TestQueries_CreateGender(t *testing.T) {
	createRandomGender(t)
}

func TestQueries_GetGender(t *testing.T) {
	g := createRandomGender(t)

	// check query with id
	g1, err := testQueries.GetGender(context.Background(), repo.GetGenderParams{ID: g.ID})
	require.NoError(t, err)
	require.Equal(t, g, g1)

	// check query with title
	g2, err := testQueries.GetGender(context.Background(), repo.GetGenderParams{Title: g.Title})
	require.NoError(t, err)
	require.Equal(t, g, g2)

	// check query without params
	g3, err := testQueries.GetGender(context.Background(), repo.GetGenderParams{})
	require.Error(t, err)
	require.Empty(t, g3)
}

func TestQueries_ListGenders(t *testing.T) {
	g1 := createRandomGender(t)
	g2 := createRandomGender(t)

	// query db
	genders, err := testQueries.ListGenders(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, genders)

	// check output
	var g1c, g2c bool
	for _, g := range genders {
		if reflect.DeepEqual(g, g1) {
			g1c = true
		} else if reflect.DeepEqual(g, g2) {
			g2c = true
		}
	}
	require.True(t, g1c, "gender g1 not selected")
	require.True(t, g2c, "gender g2 not selected")
}

func TestQueries_DeleteGender(t *testing.T) {
	g1 := createRandomGender(t)

	// delete
	err := testQueries.DeleteGender(context.Background(), g1.ID)
	require.NoError(t, err)

	// check database
	g2, err := testQueries.GetGender(context.Background(), repo.GetGenderParams{ID: g1.ID})
	require.Error(t, err)
	require.EqualError(t, sql.ErrNoRows, err.Error())
	require.Empty(t, g2)
}
