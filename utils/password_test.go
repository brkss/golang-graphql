package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

// Test Password Test hashing and verifying password mechanism
func TestPassword(t *testing.T){

	password1 := RandomString(10);
	password2 := RandomString(10);


	// valid case
	hash, err := HashPassword(password1)
	require.NoError(t, err)
	require.NotEmpty(t, hash)

	err = VerifyPassword(hash, password1)
	require.NoError(t, err)

	// invalid password
	err = VerifyPassword(hash, password2)
	require.Error(t, err)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

}
