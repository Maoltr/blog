package database

import (
	"blog/model"
	"errors"
)

func SaveArticle(title, content, username string) {
	article := model.NewArticle(title, content, username)

	db := ArticleDatabase()
	defer Close(db)

	db.Save(article)
}

func GetAllArticles() []model.Article {
	var articles []model.Article

	db := ArticleDatabase()
	defer Close(db)

	db.Find(&articles)

	return articles
}

func GetArticleById(id string) *model.Article {
	var article model.Article

	db := ArticleDatabase()
	defer Close(db)

	db.Find(&article,"id = ?", id)

	return &article
}

func DeleteArticle(id, username string) error {
	var article model.Article

	db := ArticleDatabase()
	defer Close(db)

	db.Find(&article,"id = ?", id)
	if article.Username != username {
		return errors.New("It's not your post")
	}

	db.Delete(&article)

	var res model.Article

	db.Find(&article,"id = ?", id)

	if res.ID != 0 {
		return errors.New("Can't delete article")
	}

	return nil
}

func GetUserArticles(username string) []model.Article {
	var articles []model.Article

	db := ArticleDatabase()
	defer Close(db)

	db.Find(&articles, "username = ?", username)

	return articles
}

func UpdateArticle(id, title, content, username string) string {
	var article model.Article

	db := ArticleDatabase()
	defer Close(db)

	db.Find(&article,"id = ?", id)

	if article.ID == 0 {
		return "Can't find article"
	}

	if article.Username != username {
		return "It's not your article"
	}

	db.Model(&article).Update("title", title)
	db.Model(&article).Update("content", content)

	return ""
}
