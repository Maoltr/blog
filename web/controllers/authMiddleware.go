package controllers

import (
	"github.com/gin-gonic/gin"
	. "github.com/maoltr/blog/token"
	"net/http"
)

func SetUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		userStatus(c)
	}
}

func userStatus(c *gin.Context) {
	if token, err := c.Cookie("token"); err == nil && ValidateToken(token) == nil {
		c.Set("is_logged_in", true)
	} else {
		c.Set("is_logged_in", false)
	}
}

func EnsureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedIn(c)
	}
}

func EnsureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		noLoggedIn(c)
	}
}

func loggedIn(c *gin.Context) {
	loggedInInterface, _ := c.Get("is_logged_in")
	loggedIn := loggedInInterface.(bool)

	if !loggedIn {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func noLoggedIn(c *gin.Context) {
	loggedInInterface, _ := c.Get("is_logged_in")
	loggedIn := loggedInInterface.(bool)

	if loggedIn {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
