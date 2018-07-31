package services

import (
	"fmt"
	"github.com/maoltr/blog/database"
	"github.com/maoltr/blog/model"
	"testing"
)

func TestAddUser(t *testing.T) {
	username := "Some impossible username with big quantity of letters"
	password := "password"

	err := AddUser(username, password)

	if err != nil {
		t.Error("Can't add user")
	}

	err = AddUser(username, password)

	if err == nil {
		t.Error("Problems with unique username")
	}

	DeleteUser(username)
}

func TestValidateUser(t *testing.T) {
	username := "Some impossible username with big quantity of letters"
	password := "password"

	AddUser(username, password)

	err := ValidateUser(username, password)

	if err != nil {
		t.Error("Problems with validating exist user with right password and login")
	}

	err = ValidateUser("Some impossible username with big quantity of letters but not exist", "1234")

	if err == nil {
		t.Error("Problems with validating nonexist user, he can't get access")
	}

	DeleteUser(username)
}

func TestDeleteUser(t *testing.T) {
	username := "Some impossible username with big quantity of letters"
	password := "password"

	AddUser(username, password)

	err := DeleteUser(username)

	if err != nil {
		t.Error("Problems with deleting user")
	}

	err = DeleteUser("Some impossible username with big quantity of letters but not exist")

	if err == nil {
		t.Error("You can't delete nonexist user")
	}
}

func TestUpdateUser(t *testing.T) {
	username := "Some impossible username with big quantity of letters"
	password := "password"

	AddUser(username, password)
	user := database.FindUser(username)

	id := fmt.Sprint(user.ID)
	username = "qwertyuiop[plkjhgfdsdfhjkmnbv"
	password = "sfdghjkllokijuhygtfryguhjik"

	err := UpdateUser(id, username, password)

	if err != nil {
		t.Error("Problems with updating user")
	}

	user = database.FindUser(username)

	if !compareUsers(*user, id, username, password) {
		t.Error("User wasn't update")
	}

	err = UpdateUser("1234243554", "rtyuio", "")

	if err == nil {
		t.Error("You can't update nonexist user")
	}

	user = database.FindUser("rtyuio")

	if user.ID != 0 {
		t.Error("Update method can't create users")
	}

	DeleteUser(username)
}

func compareUsers(user model.User, id, username, password string) bool {
	uId := fmt.Sprint(user.ID)

	if uId != id || username != user.Username || user.Password != password {
		return false
	}

	return true
}
