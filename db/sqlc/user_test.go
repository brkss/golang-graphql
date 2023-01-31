package db

import (
	"context"
	"testing"

	"github.com/brkss/golang-graphql/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)



func CreateRandomUser(t *testing.T) User{
	arg := RegisterUserParams{
		ID: uuid.New().String(),
		Name: utils.RandomName(),
		Email: utils.RandomEmail(),
		Password: utils.RandomString(10),
		Username: utils.RandomString(6),
	}

	user, err := testQueries.RegisterUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.ID, arg.ID)
	require.Equal(t, user.Name, arg.Name)
	require.Equal(t, user.Email, arg.Email)
	require.Equal(t, user.Password, arg.Password)
	require.Equal(t, user.Username, arg.Username)
	
	return user
}

func TestCreateUser(t *testing.T){
	CreateRandomUser(t)
}

func TestGetUser(t *testing.T) {

	_user := CreateRandomUser(t);
	
	// get user by email
	user1, err := testQueries.GetUserByIdent(context.Background(), _user.Email);
	require.NoError(t, err)
	require.NotEmpty(t, user1)

	require.Equal(t, user1.ID, _user.ID)
	require.Equal(t, user1.Name, _user.Name)
	require.Equal(t, user1.Email, _user.Email)
	require.Equal(t, user1.Username, _user.Username)

	// get user by username
	user2, err := testQueries.GetUserByIdent(context.Background(), _user.Username);
	require.NoError(t, err)
	require.NotEmpty(t, user1)

	require.Equal(t, user2.ID, _user.ID)
	require.Equal(t, user2.Name, _user.Name)
	require.Equal(t, user2.Email, _user.Email)
	require.Equal(t, user2.Username, _user.Username)
}
