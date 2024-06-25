package server

import (
	"fmt"
	jwt "pesto-auth/authorization"
	"pesto-auth/log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SignUp(c *gin.Context) {
	var user jwt.User
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
		"token":   "<enter token here>",
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
