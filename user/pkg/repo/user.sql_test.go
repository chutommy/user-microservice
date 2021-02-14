package repo_test

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
	"user/pkg/repo"
	"user/pkg/util"
)

func createRandomUser(t *testing.T) repo.User {
	t.Helper()

	now := time.Now()
	g := createRandomGender(t)

	param := repo.CreateUserParams{
		Username: sql.NullString{
			String: util.RandomString(),
			Valid:  true,
		},
		HashedPassword: util.RandomPassword(),
		FirstName:      util.RandomShortString(),
		LastName:       util.RandomLongString(),
		BirthDay: sql.NullTime{
			Time:  now.AddDate(-20, 0, 0),
			Valid: true,
		},
		Gender: g.ID,
		Email:  util.RandomEmail(),
		PhoneNumber: sql.NullString{
			String: util.RandomPhoneNumber(),
			Valid:  true,
		},
	}

	// create user
	u1, err := testQueries.CreateUser(context.Background(), param)
	require.NoError(t, err)
	require.NotEmpty(t, u1)

	// check inserted user
	u2, err := testQueries.GetUserByID(context.Background(), u1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, u2)

	// compare users' values
	compareUsers(t, u1, u2)

	return u1
}

func TestQueries_CreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestQueries_GetHashedPassword(t *testing.T) {
	u := createRandomUser(t)

	pw, err := testQueries.GetHashedPassword(context.Background(), u.ID)
	require.NoError(t, err)
	require.Equal(t, u.HashedPassword, pw) // compare password
}

func TestQueries_GetUserByID(t *testing.T) {
	u1 := createRandomUser(t)

	u2, err := testQueries.GetUserByID(context.Background(), u1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, u2)
	compareUsers(t, u1, u2)
}

func TestQueries_GetUserByUsername(t *testing.T) {
	u1 := createRandomUser(t)

	u2, err := testQueries.GetUserByUsername(context.Background(), u1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, u2)
	compareUsers(t, u1, u2)
}

func TestQueries_GetUserByEmail(t *testing.T) {
	u1 := createRandomUser(t)

	u2, err := testQueries.GetUserByEmail(context.Background(), u1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, u2)
	compareUsers(t, u1, u2)
}

// TODO: implement test cases for updates and deletions

func compareUsers(t *testing.T, u1, u2 repo.User) {
	t.Helper()

	require.Equal(t, u1.ID, u2.ID)
	require.Equal(t, u1.Username, u2.Username)
	require.Equal(t, u1.HashedPassword, u2.HashedPassword)
	require.Equal(t, u1.FirstName, u2.FirstName)
	require.Equal(t, u1.LastName, u2.LastName)
	require.Equal(t, u1.BirthDay, u2.BirthDay)
	require.Equal(t, u1.Gender, u2.Gender)
	require.Equal(t, u1.Email, u2.Email)
	require.Equal(t, u1.PhoneNumber, u2.PhoneNumber)
	require.Equal(t, u1.UpdatedAt, u2.UpdatedAt)
	require.Equal(t, u1.CreatedAt, u2.CreatedAt)
}
