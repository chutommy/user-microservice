package repo_test

import (
	"context"
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"user/pkg/repo"
	"user/pkg/util"
)

func createRandomUser(t *testing.T) repo.User {
	t.Helper()

	now := time.Now()
	g := createRandomGender(t)

	param := repo.CreateUserParams{
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

func TestQueries_GetUserByEmail(t *testing.T) {
	u1 := createRandomUser(t)

	u2, err := testQueries.GetUserByEmail(context.Background(), u1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, u2)
	compareUsers(t, u1, u2)
}

func TestQueries_UpdateUserEmail(t *testing.T) {
	u1 := createRandomUser(t)

	param := repo.UpdateUserEmailParams{
		ID:    u1.ID,
		Email: util.RandomEmail(),
	}

	u2, err := testQueries.UpdateUserEmail(context.Background(), param)
	require.NoError(t, err)
	require.NotEmpty(t, u2)

	u3, err := testQueries.GetUserByID(context.Background(), u1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, u3)

	// check values
	require.True(t, reflect.DeepEqual(u2, u3))
	require.Equal(t, u1.ID, u2.ID)
	require.Equal(t, u1.HashedPassword, u2.HashedPassword)
	require.Equal(t, u1.FirstName, u2.FirstName)
	require.Equal(t, u1.LastName, u2.LastName)
	require.Equal(t, u1.BirthDay, u2.BirthDay)
	require.Equal(t, u1.Gender, u2.Gender)
	require.Equal(t, param.Email, u2.Email)
	require.Equal(t, u1.PhoneNumber, u2.PhoneNumber)
	require.Equal(t, u1.CreatedAt, u2.CreatedAt)
	require.Equal(t, u1.DeletedAt, u2.DeletedAt)
	require.NotEqual(t, u1.UpdatedAt, u2.UpdatedAt) // was updated
}

func TestQueries_UpdateUserPassword(t *testing.T) {
	u1 := createRandomUser(t)

	param := repo.UpdateUserPasswordParams{
		ID:             u1.ID,
		HashedPassword: util.RandomPassword(),
	}

	u2, err := testQueries.UpdateUserPassword(context.Background(), param)
	require.NoError(t, err)
	require.NotEmpty(t, u2)

	u3, err := testQueries.GetUserByID(context.Background(), u1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, u3)

	// check values
	require.True(t, reflect.DeepEqual(u2, u3))
	require.Equal(t, u1.ID, u2.ID)
	require.Equal(t, param.HashedPassword, u2.HashedPassword)
	require.Equal(t, u1.FirstName, u2.FirstName)
	require.Equal(t, u1.LastName, u2.LastName)
	require.Equal(t, u1.BirthDay, u2.BirthDay)
	require.Equal(t, u1.Gender, u2.Gender)
	require.Equal(t, u1.Email, u2.Email)
	require.Equal(t, u1.PhoneNumber, u2.PhoneNumber)
	require.Equal(t, u1.DeletedAt, u2.DeletedAt)
	require.Equal(t, u1.CreatedAt, u2.CreatedAt)
	require.NotEqual(t, u1.UpdatedAt, u2.UpdatedAt) // was updated
}

func TestQueries_UpdateUserInfo(t *testing.T) {
	u0 := createRandomUser(t)
	u1 := createRandomUser(t)

	param := repo.UpdateUserInfoParams{
		ID:          u0.ID,
		FirstName:   u1.FirstName,
		LastName:    u1.LastName,
		BirthDay:    u1.BirthDay,
		Gender:      u1.Gender,
		PhoneNumber: u1.PhoneNumber,
	}

	u2, err := testQueries.UpdateUserInfo(context.Background(), param)
	require.NoError(t, err)
	require.NotEmpty(t, u2)

	u3, err := testQueries.GetUserByID(context.Background(), u0.ID)
	require.NoError(t, err)
	require.NotEmpty(t, u3)

	// check values
	require.True(t, reflect.DeepEqual(u2, u3))
	require.Equal(t, u0.ID, u2.ID)
	require.Equal(t, u1.FirstName, u2.FirstName)
	require.Equal(t, u1.LastName, u2.LastName)
	require.True(t, u1.BirthDay.Time.Equal(u2.BirthDay.Time)) // same birthday
	require.Equal(t, u1.Gender, u2.Gender)
	require.Equal(t, u1.PhoneNumber, u2.PhoneNumber)

	require.Equal(t, u0.DeletedAt, u2.DeletedAt)
	require.Equal(t, u0.CreatedAt, u2.CreatedAt)
	require.NotEqual(t, u0.UpdatedAt, u2.UpdatedAt) // was updated
}

func TestQueries_DeleteUserSoft(t *testing.T) {
	u1 := createRandomUser(t)

	// delete
	err := testQueries.DeleteUserSoft(context.Background(), u1.ID)
	require.NoError(t, err)

	// check
	u2, err := testQueries.GetUserByID(context.Background(), u1.ID)
	require.Error(t, err)
	require.EqualError(t, sql.ErrNoRows, err.Error())
	require.Empty(t, u2)

	// recover
	u3, err := testQueries.RecoverUser(context.Background(), u1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, u3)

	// check values
	require.Equal(t, u1.ID, u3.ID)
	require.Equal(t, u1.HashedPassword, u3.HashedPassword)
	require.Equal(t, u1.FirstName, u3.FirstName)
	require.Equal(t, u1.LastName, u3.LastName)
	require.Equal(t, u1.BirthDay, u3.BirthDay)
	require.Equal(t, u1.Gender, u3.Gender)
	require.Equal(t, u1.Email, u3.Email)
	require.Equal(t, u1.PhoneNumber, u3.PhoneNumber)
	require.Equal(t, u1.DeletedAt, u3.DeletedAt)
	require.Equal(t, u1.CreatedAt, u3.CreatedAt)
	require.NotEqual(t, u1.UpdatedAt, u3.UpdatedAt) // was updated
}

func TestQueries_DeleteUserPermanent(t *testing.T) {
	u1 := createRandomUser(t)

	// delete
	err := testQueries.DeleteUserPermanent(context.Background(), u1.ID)
	require.NoError(t, err)

	// check
	u2, err := testQueries.GetUserByID(context.Background(), u1.ID)
	require.Error(t, err)
	require.EqualError(t, sql.ErrNoRows, err.Error())
	require.Empty(t, u2)

	// try to recover
	u3, err := testQueries.RecoverUser(context.Background(), u1.ID)
	require.Error(t, err)
	require.EqualError(t, sql.ErrNoRows, err.Error())
	require.Empty(t, u3)
}

func compareUsers(t *testing.T, u1, u2 repo.User) {
	t.Helper()

	require.NotEmpty(t, u1)
	require.NotEmpty(t, u2)

	require.Equal(t, u1.ID, u2.ID)
	require.Equal(t, u1.HashedPassword, u2.HashedPassword)
	require.Equal(t, u1.FirstName, u2.FirstName)
	require.Equal(t, u1.LastName, u2.LastName)
	require.Equal(t, u1.BirthDay, u2.BirthDay)
	require.Equal(t, u1.Gender, u2.Gender)
	require.Equal(t, u1.Email, u2.Email)
	require.Equal(t, u1.PhoneNumber, u2.PhoneNumber)
	require.Equal(t, u1.DeletedAt, u2.DeletedAt)
	require.Equal(t, u1.UpdatedAt, u2.UpdatedAt)
	require.Equal(t, u1.CreatedAt, u2.CreatedAt)
}
