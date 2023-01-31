package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	
	return string(hash), nil
}

func VerifyPassword(hash string, password string)(error){
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}