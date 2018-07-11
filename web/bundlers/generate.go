package bundlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Render(c *gin.Context, data gin.H, templateName string) {
	loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}

func RenderErr(c *gin.Context, data gin.H, status int) {
	loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)

	switch c.Request.Header.Get("Accept") {
	case "json":
		// Respond with JSON
		c.JSON(status, data["payload"])
	case "xml":
		// Respond with XML
		c.XML(status, data["payload"])
	default:
		// Respond with HTML
		c.HTML(status, "error.html", data)
	}
}
