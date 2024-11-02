package sqlc

import (
	"context"
	"github.com/Srimanth-Reddy/simplebank/util"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	arg := CreateUserParams{
		Username: util.RandomOwner(),
		Password: hashedPassword,
		Name:     util.RandomOwner(),
		Email:    util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.UpdatedAt.IsZero())
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
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Name, user2.Name)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.UpdatedAt, user2.UpdatedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}
