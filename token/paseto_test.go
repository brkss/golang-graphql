package token

import (
	"testing"
	"time"

	"github.com/brkss/golang-graphql/utils"
	"github.com/stretchr/testify/require"
)

func TestToken(t *testing.T){

	symtricKey := utils.RandomString(32)
	maker, err := NewPasetoMaker(symtricKey)
	require.NoError(t, err)
	
	userId := utils.RandomString(10)

	// test create token !
	token, err := maker.CreateToken(userId, time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	// test varify token ! 
	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, payload.UserID, userId)
	require.WithinDuration(t, payload.ExpireAt, time.Now().Add(time.Minute), time.Second)
	require.WithinDuration(t, payload.IssuedAt, time.Now(), time.Second)
	
	// test verify token : invalid symetric key ! 
	maker1, err := NewPasetoMaker(utils.RandomString(32))
	require.NoError(t, err)
	_, err = maker1.VerifyToken(token);
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())

	// test invalid symtric key 
	_, err = NewPasetoMaker("123")
	require.Error(t, err)

}
