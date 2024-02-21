package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yangoneseok/voyager/util"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		Username: util.RandomName(),
		Email:    util.RandomEmail(),
		Password: "password",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
}
