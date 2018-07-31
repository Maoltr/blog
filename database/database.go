package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/maoltr/blog/config"
	"github.com/maoltr/blog/model"
	"sync"
)

//Path to config.sample.json
const path = "config/config.json"

var db *gorm.DB
var once sync.Once

func ArticleDatabase() *gorm.DB {
	once.Do(func() {
		conf := config.FromFile(path)
		url := conf.DB.DbConnURL()
		var err error
		fmt.Println("URL: ", url)
		//"root:w3edr509bc@/Article?parseTime=true"
		db, err = gorm.Open("mysql", url)

		db.AutoMigrate(&model.User{})
		db.AutoMigrate(&model.Article{})

		if err != nil {
			panic("Failed to connect database: %s" + err.Error())
		}
	})

	return db
}

func Close(db *gorm.DB) {
	err := db.Close()

	if err != nil {
		panic("Can't close database")
	}
}

