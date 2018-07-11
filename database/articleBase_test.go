package database

import (
	"blog/model"
	"fmt"
	"testing"
)

func TestArticleBase(t *testing.T) {
	username := "someNewUserWith"
	title := "TestArticle"
	content := "TestContentForTestArticle"

	SaveArticle(title, content, username)

	articles := GetUserArticles(username)

	//Checking that article was save and we can get it by GetUserArticles
	if len(articles) <= 0 && !compareArticles(articles[0], title, content, username) {
		t.Error("No articles or invalid article")
	}

	id := fmt.Sprint(articles[0].ID)

	DeleteArticle(id, username)

	article := GetArticleById(id)

	//Checking that article was delete and we can't get it by GetArticleById
	if article.ID != 0 {
		t.Error("Article wasn't deleted or GetArticleById doesn't work")
	}

	SaveArticle(title, content, username)

	//Checking that we can get Articles by GetAllArticles
	articles = GetAllArticles()
	userArticles := GetUserArticles(username)

	if len(articles) <= 0 {
		t.Error("Can't get articles from base")
	}

	id = fmt.Sprint(userArticles[0].ID)

	//Checking that we can't delete stranger article
	err := DeleteArticle(id, "Joni")

	if err == nil {
		t.Error("You can't delete not your article")
	}

	//Clear our base
	DeleteArticle(id, username)
}

//Compare article from base and our data
func compareArticles(article model.Article, title, content, username string) bool {
	if article.Content != content || article.Username != username || article.Title != title {
		return false
	}

	return true
}
