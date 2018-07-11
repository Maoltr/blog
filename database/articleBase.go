package database

import (
	"blog/model"
	"errors"
	"fmt"
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

	db.First(&article, id)

	return &article
}

func DeleteArticle(id, username string) error {
	var article model.Article

	db := ArticleDatabase()
	defer Close(db)
	fmt.Println("Try to find post: ", id)
	db.First(&article, id)
	if article.Username != username {
		return errors.New("It's not your post")
	}
	fmt.Println("Deleting")
	db.Delete(&article)

	var res model.Article

	db.First(&res, id)

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
