package jwt

import (
	"pesto-auth/log"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Logger.Debug("Hash check", zap.String("hash", hash), zap.String("password", password))
		log.Logger.Error("Failed comparing Hash and Password :", zap.Error(err))
	}
	return err == nil
}
