package services

import (
	"testing"
	"fmt"
)

func TestServiceArticle(t *testing.T) {
	title := "test article"
	content := "some content for test article"
	username := "asdfghjklfghjkjhghjdfsdfdsfsd"

	SaveArticle(title, content, username)

	articles := GetUserArticles(username)

	if len(articles) <= 0 {
		t.Error("Problems with save or getUserArticles")
	}

	id := fmt.Sprint(articles[0].Id)

	DeleteArticle(id, username)

	article := GetArticle(id)

	if article.Id != 0 {
		t.Error("Problems with deleting article")
	}

	SaveArticle(title, content, username)

	articles = GetAllArticles()
	arts := GetUserArticles(username)
	id = fmt.Sprint(arts[0].Id)

	if len(articles) <= 0 {
		t.Error("Problems with GetAllArticles")
	}

	DeleteArticle(id, username)
}
