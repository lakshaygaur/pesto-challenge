package jwt

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		fmt.Println("hash :", hash, " password : ", password)
		fmt.Println("Failed comparing Hash and Password :", err)
	}
	return err == nil
}
