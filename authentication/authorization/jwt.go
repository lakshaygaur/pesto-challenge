package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

type User struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Country string `json:"country"`
	Phone   string `json:"phone"`
	jwt.StandardClaims
}

var secretKey []byte
var EXPIRY int

func Init(cfg Config) {
	secretKey = []byte(cfg.Secret)
	EXPIRY = cfg.Expiry
}

func CreateToken(user User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		user)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Verify(tokenString string) (*User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &User{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	user := token.Claims.(*User)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return user, nil
}
