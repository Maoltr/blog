package controllers

import (
	"blog/config"
	"blog/model"
	"blog/services"
	"blog/web/bundlers"
	"github.com/gin-gonic/gin"
	"net/http"
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
		bundlers.Render(c, gin.H{
			"title":   "Article",
			"payload": article,
		}, "article.html")
		return
	}

	loggedInInterface, _ := c.Get("is_logged_in")

	c.HTML(http.StatusNotFound, "article.html", gin.H{
		"title":        "Article",
		"payload":      model.NewArticle("Not found", "", ""),
		"is_logged_in": loggedInInterface.(bool),
	})
}

func DeleteArticle(c *gin.Context) {
	username := getUsername(c)

	id := c.Param("id")

	er := services.DeleteArticle(id, username)

	if er == nil {
		GetAllArticles(c)
	}
}

func ManageArticles(c *gin.Context) {
	username := getUsername(c)
	loggedInInterface, _ := c.Get("is_logged_in")

	articles := services.GetUserArticles(username)

	if len(articles) <= 0 {
		c.HTML(http.StatusNotFound, "error.html", gin.H{
			"title":        "Error",
			"text":         "You haven't articles, create them:)",
			"is_logged_in": loggedInInterface.(bool),
		})
		return
	}

	bundlers.Render(c, gin.H{
		"title":   "My Articles",
		"payload": articles,
	}, "manageArticles.html")
}

func notAuthorized(c *gin.Context) {
	loggedInInterface, _ := c.Get("is_logged_in")

	c.HTML(http.StatusUnauthorized, "home.html", gin.H{
		"title":        "Not authorized",
		"is_logged_in": loggedInInterface.(bool),
	})
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
