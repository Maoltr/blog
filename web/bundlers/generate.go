package bundlers

import (
	"github.com/gin-gonic/gin"
	"html/template"
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
	case "article":
		tmpl := template.Must(template.ParseFiles("/home/maksim/go/src/blog/resources/templates/article.html",
			"/home/maksim/go/src/blog/resources/templates/footer.html", "/home/maksim/go/src/blog/resources/templates/header.html"))
		err := tmpl.Execute(c.Writer, data)

		if !checkErr(err) {
			RenderErr(c, gin.H{
				"title": "Article",
				"text":  err.Error(),
			}, http.StatusBadRequest)
		}
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}

func RenderErr(c *gin.Context, data gin.H, status int) {
	loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(status, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(status, data["payload"])
	default:
		// Respond with HTML
		c.HTML(status, "error.html", data)
	}
}

func checkErr(err error) bool {
	if err != nil {
		return false
	} else {
		return true
	}
}
