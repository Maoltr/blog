package controllers

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/gin-gonic/gin"
	"blog/services"
	"fmt"
	"net/url"
	"blog/config"
	"time"
)

var (
	username = "userdfgdfg"
	title = "title"
	content = "content"
)

func TestGetAllArticles(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/", nil)

	c.Set("is_logged_in", true)
	c.Request.Header.Set("Accept", "application/json")

	GetAllArticles(c)

	if w.Code != http.StatusOK {
		t.Error("Status: ", w.Code, " Body: ", w.Body)
	}
}

func TestGetArticle(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	id := createArticle()

	c.Request, _ = http.NewRequest("GET", "/article/view/" + id, nil)

	c.Set("is_logged_in", true)
	c.Request.Header.Set("Accept", "application/json")

	GetArticle(c)

	if w.Code != http.StatusOK {
		t.Error("Status: ", w.Code, " Body: ", w.Body)
	}

	services.DeleteArticle(id, username)
}

func TestCreateArticle(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("GET", "/article/create/", nil)

	c.Set("is_logged_in", true)
	c.Request.Header.Set("Accept", "application/json")

	CreateArticle(c)

	if w.Code != http.StatusOK {
		t.Error("Status: ", w.Code, " Body: ", w.Body)
	}
}

func TestPostAndDeleteArticle(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	prepareContext(c)

	PostArticle(c)

	if w.Code != http.StatusOK {
		t.Error("Status: ", w.Code, " Body: ", w.Body)
	}

	article := services.GetUserArticles("testuser")[0]
	id := fmt.Sprint(article.Id)

	c.Request.URL.Path += "/" + id
	DeleteArticle(c)

	if w.Code != http.StatusOK {
		t.Error("Can't delete article, code: ", w.Code)
	}
}

func TestManageArticles(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	id := createArticle()

	c.Request, _ = http.NewRequest("GET", "/article/manage/", nil)
	c.Set("is_logged_in", true)
	c.Request.Header.Set("Accept", "application/json")
	token := config.GenerateToken(username, time.Duration(time.Second * 5))
	c.Request.Header.Set("Cookie", "token="+token)

	ManageArticles(c)

	services.DeleteArticle(id, username)

	if w.Code != http.StatusOK {
		t.Error("Error code: ", w.Code)
	}
}

func createArticle() string{
	services.SaveArticle(title, content, username)

	article := services.GetUserArticles(username)[0]
	id := fmt.Sprint(article.Id)

	return id
}

func prepareContext(c *gin.Context) {
	c.Request, _ = http.NewRequest("POST", "/article/create/", nil)

	c.Set("is_logged_in", true)
	c.Request.Header.Set("Accept", "application/json")
	c.Request.PostForm = url.Values{}
	c.Request.PostForm.Set("title", "title")
	c.Request.PostForm.Set("content", "content")
	token := config.GenerateToken("testuser", time.Duration(time.Second * 5))
	c.Request.Header.Set("Cookie", "token="+token)
}


