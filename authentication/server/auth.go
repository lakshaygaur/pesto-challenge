package server

import (
	"fmt"
	jwt "pesto-auth/authorization"
	"pesto-auth/log"
	"pesto-auth/user"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func SignUp(c *gin.Context) {
	pass, err := jwt.GenerateHashPassword("mystrongpass")
	if err != nil {
		log.Logger.Error("Failed generating password : ", zap.Error(err))
	}
	var user = user.User{
		Id:       uuid.NewString(),
		Name:     "Test",
		Email:    "test@example.com",
		Password: pass,
		Country:  "India",
		Phone:    "+91-989749834",
	}
	fmt.Println("user in auth ", user)
	user.StoreUser()
	// var user user.User
	if c.ShouldBind(&user) != nil {
		log.Logger.Error("bind failed")
	}
	log.Logger.Debug("User details for signup ", zap.Any("user", user))
	token, err := jwt.CreateToken(user)
	if err != nil {
		log.Logger.Error("Failed creating token : ", zap.String("error", err.Error()))
	}
	// log.Logger.Debug("Token Created : ", token)
	fmt.Println("Token Created : ", token)
	c.JSON(200, gin.H{
		"success": true,
		"token":   token,
	})
}

func Login(c *gin.Context) {
	// var token string
	// if c.ShouldBind(&string) != nil {
	// 	log.Logger.Error("bind failed")
	// }
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiVGVzdCIsImVtYWlsIjoidGVzdEBleGFtcGxlLmNvbSIsImNvdW50cnkiOiJJbmRpYSIsInBob25lIjoiSW5kaWEiLCJleHAiOjE3MTkyNTcyMjksImlhdCI6MTcxOTI1NjMyOX0.D5UfmtkoUjGTyjdIBcWMcYCHrajpdTu1mEHcTLSQzGI"
	jwt.Verify(token)
	c.JSON(200, gin.H{
		"success": true,
		"token":   "<enter token here>",
	})
}

func User(c *gin.Context) {

	c.JSON(200, gin.H{
		"success": true,
		"token":   "<enter token here>",
	})
}
