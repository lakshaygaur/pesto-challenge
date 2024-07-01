package user

import (
	"errors"
	"pesto-auth/database"
	"pesto-auth/log"

	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Country  string `json:"country"`
	Phone    string `json:"phone"`
	jwt.StandardClaims
}

func (user *User) StoreUser() error {
	// perform validations
	// check email
	// check phone
	// insert finally
	data, err := database.DB.Exec(`INSERT INTO users ( id, name, email, password, country, phone)
	 VALUES ( $1, $2, $3, $4, $5, $6)`, user.Id, user.Name, user.Email, user.Password, user.Country, user.Phone)
	if err != nil {
		log.Logger.Error("Failed storing user in DB : ", zap.Error(err))
		return err
	}
	log.Logger.Debug("User inserted with details : ", zap.Any("user", data))
	return nil
}

func (user *User) GetUser() error {
	if user.Email == "" {
		return errors.New("User object has no email")
	}
	data, err := database.DB.Query("select * from users where email=$1", user.Email)
	if err != nil {
		log.Logger.Error("Failed getting user details : ", zap.Error(err))
		return err
	}
	// set user
	for data.Next() {
		err = data.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Country, &user.Phone)
		if err != nil {
			log.Logger.Error("Failed scanning user details : ", zap.Error(err))
			return err
		}
	}
	log.Logger.Debug("User details scanned : ", zap.Any("user", user))
	return nil
}
