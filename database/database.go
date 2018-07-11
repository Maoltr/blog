package database

import (
	"blog/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Database() *gorm.DB {
	db, err := gorm.Open("mysql", "root:w3edr509bc@/Product?parseTime=true")

	if err != nil {
		panic("Failed to connect database: %s" + err.Error())
	}

	return db
}

func ArticleDatabase() *gorm.DB {
	db, err := gorm.Open("mysql", "root:w3edr509bc@/Article?parseTime=true")
	db.AutoMigrate(&model.Article{})
	db.AutoMigrate(&model.User{})
	if err != nil {
		panic("Failed to connect database: %s" + err.Error())
	}

	return db
}

func Close(db *gorm.DB) {
	err := db.Close()

	if err != nil {
		panic("Can't close database")
	}
}
