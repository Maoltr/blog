package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title    string `json:"title"`
	Content  string `json:"content"`
	Username string `json:"username"`
}

func NewArticle(title, content, username string) *Article {
	article := Article{
		Title:    title,
		Content:  content,
		Username: username,
	}

	return &article
}
