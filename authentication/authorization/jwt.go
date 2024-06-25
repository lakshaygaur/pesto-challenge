package jwt

import (
	"fmt"
	"pesto-auth/user"

	"github.com/golang-jwt/jwt"
)

var secretKey []byte
var EXPIRY int

func Init(cfg Config) {
	secretKey = []byte(cfg.Secret)
	EXPIRY = cfg.Expiry
}

func CreateToken(user user.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		user)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Verify(tokenString string) (*user.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &user.User{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	user := token.Claims.(*user.User)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return user, nil
}
