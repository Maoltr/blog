package services

import (
	"github.com/maoltr/blog/database"
	"github.com/maoltr/blog/model"
)

func GetAllArticles() []model.TransformedArticle {
	articles := database.GetAllArticles()
	var res []model.TransformedArticle

	for _, item := range articles {
		res = append(res, *model.NewTransformedArticle(item))
	}

	return res
}

func SaveArticle(title, content, username string) {
	database.SaveArticle(title, content, username)
}

func GetArticle(id string) *model.TransformedArticle {
	article := database.GetArticleById(id)

	res := model.NewTransformedArticle(*article)

	return res
}

func DeleteArticle(id, username string) error {
	err := database.DeleteArticle(id, username)

	return err
}

func GetUserArticles(username string) []model.TransformedArticle {
	articles := database.GetUserArticles(username)
	var res []model.TransformedArticle

	for _, item := range articles {
		res = append(res, *model.NewTransformedArticle(item))
	}

	return res
}

func UpdateArticle(id, title, content, username string) string {
	err := database.UpdateArticle(id, title, content, username)
	return err
}
