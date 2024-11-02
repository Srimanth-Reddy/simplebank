package api

import (
	"github.com/Srimanth-Reddy/simplebank/db/sqlc"
	"github.com/Srimanth-Reddy/simplebank/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func randomUser(t *testing.T) (user sqlc.User, password string) {
	password = util.RandomString(6)
	hashedPassword, err := util.HashPassword(password)
	require.NoError(t, err)

	user = sqlc.User{
		Username: util.RandomOwner(),
		Password: hashedPassword,
		Name:     util.RandomOwner(),
		Email:    util.RandomEmail(),
	}
	return
}
