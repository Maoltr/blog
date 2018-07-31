package database

import (
	"github.com/maoltr/blog/model"
)

func SaveUser(username, password string) bool {
	user := model.NewUser(username, password)

	if IsUsernameAvailable(username) {
		db := ArticleDatabase()
		db.Save(&user)

		return true
	}

	return false
}

func IsUsernameAvailable(username string) bool {
	var user model.User

	db := ArticleDatabase()
	db.First(&user, "username = ?", username)

	if user.ID != 0 {
		return false
	}

	return true
}

func FindUser(username string) *model.User {
	var user model.User

	db := ArticleDatabase()
	db.Find(&user, "username = ?", username)

	return &user
}

func DeleteUser(username string) bool {
	var user model.User

	db := ArticleDatabase()
	db.Find(&user, "username = ?", username)

	if user.ID != 0 {
		db.Delete(&user)
		return true
	}

	return false
}

func UpdateUser(id, username, password string) *model.User {
	var user model.User

	db := ArticleDatabase()
	db.Find(&user, id)

	if user.ID != 0 {
		db.Model(&user).Update("username", username)
		db.Model(&user).Update("password", password)
	}

	return &user
}
