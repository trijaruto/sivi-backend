package security

import (
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

func ComparePassword(dbpassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbpassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}
