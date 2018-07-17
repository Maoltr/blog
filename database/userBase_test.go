package database

import (
	"fmt"
	"testing"
)

func TestSaveAndFindAndDeleteUser(t *testing.T) {
	username := "SomeNewUserWithSmt"
	password := "12345678"

	//Save our user
	SaveUser(username, password)

	//Try to find saved user
	user := FindUser(username)

	//Checking that our user was save
	if user.Username != username || user.Password != password || user.ID == 0 {
		t.Error("Can't find or save user")
	}

	//Checking that we can delete him
	if !DeleteUser(username) {
		t.Error("Can't delete user")
	}

	//Checking that we can't delete nonexist user
	if DeleteUser("dgsgfdsgdfgdfghdfsgfdsg") {
		t.Error("Deleted nonexist user")
	}

	//Checking that we can't find nonexist user
	user = FindUser("sgfsgdhdfjgfjsdfsdfsfgf")

	if user.ID != 0 {
		t.Error("Found nonexist user")
	}
}

func TestIsUsernameAvailable(t *testing.T) {
	username := "SomeNewUserWithSmt"
	password := "12345678"

	SaveUser(username, password)

	//Check that we name, which we use for our user non available
	if IsUsernameAvailable(username) {
		t.Error("This name is busy, but IsUsernameAvailable tell that it's free")
	}

	//Check that random unpronounceable name is free
	if !IsUsernameAvailable("fsdhgdjkhfopkhghjgkhgkhgsdfgdfsgdfk") {
		t.Error("This name is free, but IsUsernameAvailable tell that it's busy")
	}

	//Clear base
	DeleteUser(username)
}

func TestUpdateUser(t *testing.T) {
	username := "SomeNewUserWithSmt"
	password := "12345678"

	SaveUser(username, password)
	//Get our user from base
	user := FindUser(username)
	id := fmt.Sprint(user.ID)

	//Data for updating our user
	username = "SomeNewUserWithSmt2"
	password = "87654321"

	user = UpdateUser(id, username, password)

	//Checking that user was update
	if user.Username != username || user.Password != password {
		t.Error("Faild update user")
	}

	//Clear base
	DeleteUser(username)
}
