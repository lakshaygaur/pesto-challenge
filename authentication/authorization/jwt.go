package jwt

import (
	"encoding/json"
	"pesto-auth/user"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

var secretKey []byte
var EXPIRY int

func Init(cfg Config) {
	secretKey = []byte(cfg.Secret)
	EXPIRY = cfg.Expiry
}

func CreateToken(user user.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user": user,
			"exp":  (time.Now().Add(10 * time.Minute)).UnixMilli(),
		})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	// refreshClaims := jwt.StandardClaims{
	// 	IssuedAt:  time.Now().Unix(),
	// 	ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
	//   }

	//   signedRefreshToken, err := service.NewRefreshToken(refreshClaims)
	//   if err != nil {
	// 	log.Fatal("error creating refresh token")
	//   }

	return tokenString, nil
}

func Verify(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	parsed := token.Claims.(jwt.MapClaims)
	if err != nil {
		return nil, err
	}
	var exp int64 //time.Time
	jsonbody, err := json.Marshal(parsed["exp"])
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonbody, &exp)
	if err != nil {
		return nil, err
	}
	expiryTime := time.UnixMilli(exp)
	if expiryTime.UnixMilli() < time.Now().UnixMilli() {
		return nil, errors.New("token expired")
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return parsed, nil
}
