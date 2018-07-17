package model

import (
	"time"
)

type Article struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
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
