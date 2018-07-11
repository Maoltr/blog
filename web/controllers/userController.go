package controllers

import (
	"blog/config"
	"blog/services"
	"blog/web/bundlers"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func ShowLoginPage(c *gin.Context) {
	bundlers.Render(c, gin.H{
		"title": "Login",
	}, "login.html")
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	err := services.ValidateUser(username, password)

	if err == nil {
		c.SetCookie("token", config.GenerateToken(username, time.Duration(time.Minute*20)), 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		bundlers.Render(c, gin.H{
			"title": "Successful login",
		}, "login-successful.html")

		return
	}

	c.HTML(http.StatusBadRequest, "login.html", gin.H{
		"ErrorTitle":   "Login Failed",
		"ErrorMessage": err.Error(),
	})
}

func LogOut(c *gin.Context) {
	c.SetCookie("token", "", -1, "", "", false, true)
	GetAllArticles(c)
}

func ShowRegisterPage(c *gin.Context) {
	bundlers.Render(c, gin.H{
		"title": "Registration",
	}, "register.html")
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	err := services.AddUser(username, password)

	if err == nil {
		c.SetCookie("token", config.GenerateToken(username, time.Duration(time.Minute*20)), 3600, "", "", false, true)
		c.Set("is_logged_in", true)

		bundlers.Render(c, gin.H{
			"title": "Successful registration",
		}, "login-successful.html")
		return
	}

	loggedInInterface, _ := c.Get("is_logged_in")

	c.HTML(http.StatusBadRequest, "login.html", gin.H{
		"ErrorTitle":   "Registration Failed",
		"ErrorMessage": err.Error(),
		"is_logged_in": loggedInInterface.(bool),
	})
}
