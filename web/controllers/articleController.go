package controllers

import (
	"blog/config"
	"blog/services"
	"blog/web/bundlers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GetAllArticles(c *gin.Context) {
	data := services.GetAllArticles()

	bundlers.Render(c, gin.H{
		"title":   "Home Page",
		"payload": data,
	}, "index.html")

}

func PostArticle(c *gin.Context) {
	username := getUsername(c)

	services.SaveArticle(c.PostForm("title"), c.PostForm("content"), username)

	GetAllArticles(c)
}

func CreateArticle(c *gin.Context) {
	bundlers.Render(c, gin.H{
		"title": "Create Article",
	}, "create-article.html")
}

func GetArticle(c *gin.Context) {
	article := services.GetArticle(c.Param("id"))

	if article.Id != 0 {
		article.Content = strings.Replace(article.Content, "\n", "<br>", -1)

		c.Request.Header.Set("Accept", "article")
		bundlers.Render(c, gin.H{
			"title":   "Article",
			"payload": article,
		}, "article.html")

		c.Request.Header.Set("Accept", "")

		return
	}

	bundlers.RenderErr(c, gin.H{
		"title":        "Article",
		"text":         "Not found",
	}, http.StatusNotFound)

}

func DeleteArticle(c *gin.Context) {
	username := getUsername(c)

	id := c.Param("id")

	er := services.DeleteArticle(id, username)

	if er == nil {
		GetAllArticles(c)
		return
	}

	bundlers.RenderErr(c, gin.H{
		"title":        "Deleting",
		"text":         "Can't delete article, try again later",
	}, http.StatusBadRequest)
}

func ManageArticles(c *gin.Context) {
	username := getUsername(c)

	articles := services.GetUserArticles(username)

	if len(articles) <= 0 {
		bundlers.RenderErr(c, gin.H{
			"title":        "You haven't articles",
			"text":         "You haven't articles, create them:)",
		}, http.StatusNotFound)
		return
	}

	bundlers.Render(c, gin.H{
		"title":   "My Articles",
		"payload": articles,
	}, "manageArticles.html")
}

func GetUpdateArticle(c *gin.Context) {
	id := c.Param("id")

	article := services.GetArticle(id)

	if article.Id == 0 {
		bundlers.RenderErr(c, gin.H{
			"title": "Not found article",
			"text": "Can't found this article",
		}, http.StatusNotFound)
		return
	}

	bundlers.Render(c, gin.H{
		"title": "Article",
		"payload": article,
	}, "update-article.html")
}

func PostUpdateArticle(c *gin.Context) {
	id := c.Param("id")

	title := c.PostForm("title")
	content := c.PostForm("content")
	username := getUsername(c)

	err := services.UpdateArticle(id, title, content, username)

	if err == "" {
		GetAllArticles(c)
		return
	}
	bundlers.RenderErr(c, gin.H{
		"title": "Can't update article",
		"text": err,
	}, http.StatusBadRequest)

	return

}

func notAuthorized(c *gin.Context) {
	bundlers.RenderErr(c, gin.H{
		"title":        "Not authorized",
		"text":         "You don't authorized",
	}, http.StatusUnauthorized)
}

func getUsername(c *gin.Context) string {
	token, err := c.Cookie("token")

	if err != nil {
		notAuthorized(c)
		return ""
	}

	username, err := config.GetUsername(token)

	if err != nil {
		notAuthorized(c)
		return ""
	}

	return username
}
