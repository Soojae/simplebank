package db

import (
	"context"
	"database/sql"
	"github.com/soojae/simplebank/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err2 := util.HashPassword(util.RandomString(6))
	require.NoError(t, err2)

	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)
	return user
}

func TestCreateUser(t *testing.T) {

	createRandomUser(t)

}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUserOnlyFullName(t *testing.T) {
	oldUser := createRandomUser(t)

	newFullName := util.RandomOwner()
	arg := UpdateUserParams{
		Username: oldUser.Username,
		FullName: sql.NullString{
			String: newFullName,
			Valid:  true,
		},
	}

	updatedUser, err := testQueries.UpdateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEqual(t, oldUser.FullName, updatedUser.FullName)
	require.Equal(t, newFullName, updatedUser.FullName)
	require.Equal(t, oldUser.HashedPassword, updatedUser.HashedPassword)
	require.Equal(t, oldUser.Email, updatedUser.Email)
	require.WithinDuration(t, oldUser.PasswordChangedAt, updatedUser.PasswordChangedAt, time.Second)
	require.WithinDuration(t, oldUser.CreatedAt, updatedUser.CreatedAt, time.Second)
}

func TestUpdateUserOnlyEmail(t *testing.T) {
	oldUser := createRandomUser(t)

	newEmail := util.RandomEmail()
	arg := UpdateUserParams{
		Username: oldUser.Username,
		Email: sql.NullString{
			String: newEmail,
			Valid:  true,
		},
	}

	updatedUser, err := testQueries.UpdateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEqual(t, oldUser.Email, updatedUser.Email)
	require.Equal(t, newEmail, updatedUser.Email)
	require.Equal(t, oldUser.FullName, updatedUser.FullName)
	require.Equal(t, oldUser.HashedPassword, updatedUser.HashedPassword)
	require.WithinDuration(t, oldUser.PasswordChangedAt, updatedUser.PasswordChangedAt, time.Second)
	require.WithinDuration(t, oldUser.CreatedAt, updatedUser.CreatedAt, time.Second)
}

func TestUpdateUserOnlyPassword(t *testing.T) {
	oldUser := createRandomUser(t)

	newPassword := util.RandomString(6)
	newHashedPassword, err := util.HashPassword(newPassword)
	require.NoError(t, err)

	arg := UpdateUserParams{
		Username: oldUser.Username,
		HashedPassword: sql.NullString{
			String: newHashedPassword,
			Valid:  true,
		},
	}

	updatedUser, err := testQueries.UpdateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEqual(t, oldUser.HashedPassword, updatedUser.HashedPassword)
	require.Equal(t, newHashedPassword, updatedUser.HashedPassword)
	require.Equal(t, oldUser.FullName, updatedUser.FullName)
	require.Equal(t, oldUser.Email, updatedUser.Email)
	require.WithinDuration(t, oldUser.PasswordChangedAt, updatedUser.PasswordChangedAt, time.Second)
	require.WithinDuration(t, oldUser.CreatedAt, updatedUser.CreatedAt, time.Second)
}

func TestUpdatedUserAllFields(t *testing.T) {

	oldUser := createRandomUser(t)

	newFullName := util.RandomOwner()
	newEmail := util.RandomEmail()
	newPassword := util.RandomString(6)
	newHashedPassword, err := util.HashPassword(newPassword)
	require.NoError(t, err)

	arg := UpdateUserParams{
		Username: oldUser.Username,
		FullName: sql.NullString{
			String: newFullName,
			Valid:  true,
		},
		Email: sql.NullString{
			String: newEmail,
			Valid:  true,
		},
		HashedPassword: sql.NullString{
			String: newHashedPassword,
			Valid:  true,
		},
	}

	updatedUser, err := testQueries.UpdateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEqual(t, oldUser.FullName, updatedUser.FullName)
	require.NotEqual(t, oldUser.Email, updatedUser.Email)
	require.NotEqual(t, oldUser.HashedPassword, updatedUser.HashedPassword)
	require.Equal(t, newFullName, updatedUser.FullName)
	require.Equal(t, newEmail, updatedUser.Email)
	require.Equal(t, newHashedPassword, updatedUser.HashedPassword)
	require.WithinDuration(t, oldUser.PasswordChangedAt, updatedUser.PasswordChangedAt, time.Second)
	require.WithinDuration(t, oldUser.CreatedAt, updatedUser.CreatedAt, time.Second)

}
