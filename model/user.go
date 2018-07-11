package model

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUser(username, password string) *User {
	user := User{
		Username: username,
		Password: password,
	}

	return &user
}

func Validate(user User, username, password string) error {
	if user.Password == password && user.Username == username {
		return nil
	}

	return errors.New("Bad password or username")
}
