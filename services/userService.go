package services

import (
	"github.com/maoltr/blog/database"
	"github.com/maoltr/blog/model"
	"github.com/pkg/errors"
)

func ValidateUser(username, password string) error {
	var user model.User

	user = *database.FindUser(username)

	err := model.Validate(user, username, password)

	if err == nil {
		return nil
	}

	return err
}

func AddUser(username, password string) error {
	if database.SaveUser(username, password) {
		return nil
	}

	return errors.New("Please change username, this is busy")
}

func DeleteUser(username string) error {
	if database.DeleteUser(username) {
		return nil
	}

	return errors.New("Can't delete user")
}

func UpdateUser(id, username, password string) error {
	if database.UpdateUser(id, username, password).ID != 0 {
		return nil
	}

	return errors.New("Can't update user")
}
