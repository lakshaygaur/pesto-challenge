package server

import (
	"encoding/json"
	"net/http"
	jwt "pesto-auth/authorization"
	"pesto-auth/log"
	"pesto-auth/user"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func SignUp(c *gin.Context) {
	var user user.User
	if c.ShouldBind(&user) != nil {
		log.Logger.Error("failed binding user : ")
		return
	}
	pass, err := jwt.GenerateHashPassword(user.Password)
	if err != nil {
		log.Logger.Error("Failed generating password : ", zap.Error(err))
		return
	}
	user.Password = pass
	user.Id = uuid.NewString()
	log.Logger.Debug("User details for signup ", zap.Any("user", user))
	token, err := jwt.CreateToken(user)
	if err != nil {
		log.Logger.Error("Failed creating token : ", zap.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"msg":     "Failed adding user details. Please contact support.",
		})
		return
	}
	err = user.StoreUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"msg":     "Failed adding user details. Please contact support.",
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"token":   token,
	})
}

func Login(c *gin.Context) {
	var reqUsr user.User
	if c.ShouldBind(&reqUsr) != nil {
		log.Logger.Error("bind failed")
	}
	// match passwords
	tempPass := reqUsr.Password
	reqUsr.GetUser()
	ok := jwt.CheckPasswordHash(tempPass, reqUsr.Password)
	if !ok {
		c.JSON(200, gin.H{
			"success": false,
			"msg":     "Invalid email/password",
		})
		return
	}
	// generate token
	token, err := jwt.CreateToken(reqUsr)
	if err != nil {
		log.Logger.Error("Failed creating token : ", zap.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"msg":     "Failed adding user details. Please contact support.",
		})
		return
	}
	// return token
	jwt.Verify(token)
	c.JSON(200, gin.H{
		"success": true,
		"token":   token,
	})
}

func RefreshToken(c *gin.Context) {
	// parse token

	// check if the token has expired; reject request if it is expired

	// generate new token with more time

	c.JSON(200, gin.H{
		"success": true,
		"token":   "<>",
	})
}

func User(c *gin.Context) {
	// check if token is valid
	authorization := c.GetHeader("Authorization")
	// Check if the Authorization header has the correct format (Bearer <token>)
	token := strings.Split(authorization, " ")
	if len(token) != 2 || token[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
		c.Abort()
		return
	}
	tokenMap, err := jwt.Verify(token[1])
	if err != nil {
		log.Logger.Error("Failed verifying token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"msg":     "Failed verifying token",
		})
		c.Abort()
		return
	}

	jsonbody, err := json.Marshal(tokenMap["user"])
	if err != nil {
		log.Logger.Error("Failed parsing token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"msg":     "Failed parsing token",
		})
		c.Abort()
		return
	}
	user := user.User{}
	if err := json.Unmarshal(jsonbody, &user); err != nil {
		log.Logger.Error("Failed parsing token", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"msg":     "Failed parsing token",
		})
		c.Abort()
		return
	}

	// get details from DB
	err = user.GetUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"msg":     "Failed getting user details",
		})
		c.Abort()
		return
	}
	// build response object
	resp := make(map[string]string)
	resp["name"] = user.Name
	resp["email"] = user.Email
	resp["id"] = user.Id
	resp["country"] = user.Country
	resp["phone"] = user.Phone

	c.JSON(200, gin.H{
		"success": true,
		"user":    resp,
	})
}
