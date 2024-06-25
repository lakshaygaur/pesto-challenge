package jwt

import (
	"fmt"
	"pesto-auth/user"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
)

func TestCreateToken(t *testing.T) {
	var user = user.User{
		Name:    "Test",
		Email:   "test@example.com",
		Country: "India",
		Phone:   "India",
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	t.Run("Create and verify token", func(t *testing.T) {
		token, err := CreateToken(user)
		if err != nil {
			t.Errorf("Failed creating token %q", err)
		}
		fmt.Println("Token created : ", token)
		got, err := Verify(token)
		if err != nil {
			t.Errorf("Failed verifying token %q", err)
		}
		want := "Test"
		if got.Name != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
}
