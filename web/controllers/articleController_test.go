package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/maoltr/blog/services"
	. "github.com/maoltr/blog/token"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

var (
	username = "userdfgdfg"
	title    = "title"
	content  = "content"
)

func TestGetAllArticles(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	//Prepare our context
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
	//Create article which we want get
	id := createArticle()
	//Prepare our context
	c.Request, _ = http.NewRequest("GET", "/article/view", nil)
	params := make([]gin.Param, 10)
	c.Params = params
	c.Params[0] = gin.Param{Key: "id", Value: id}
	c.Set("is_logged_in", true)
	c.Request.Header.Set("Accept", "application/json")

	GetArticle(c)

	if w.Code != http.StatusOK {
		t.Error("Status: ", w.Code, " Body: ", w.Body)
	}

	services.DeleteArticle(id, username)

	//Checking that we can't get nonexist article
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("GET", "/article/view/"+id, nil)

	c.Set("is_logged_in", true)
	c.Request.Header.Set("Accept", "application/json")

	GetArticle(c)

	if w.Code != http.StatusNotFound {
		t.Error("Status: ", w.Code, " Body: ", w.Body)
	}
}

func TestCreateArticle(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	//Prepare our context
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

	//Checking that we can post article
	PostArticle(c)

	if w.Code != http.StatusOK {
		t.Error("Status: ", w.Code, " Body: ", w.Body)
	}

	article := services.GetUserArticles("testuser")[0]
	id := fmt.Sprint(article.Id)

	//Checking that we can delete exist post
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	prepareContext(c)
	params := make([]gin.Param, 10)
	c.Params = params
	c.Params[0] = gin.Param{Key: "id", Value: id}

	DeleteArticle(c)

	if w.Code != http.StatusOK {
		t.Error("Can't delete article, code: ", w.Code)
	}

	//Checking that we can't delete non exist post
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	prepareContext(c)
	c.Request.URL.Path += "/" + id

	DeleteArticle(c)

	if w.Code != http.StatusBadRequest {
		t.Error("Can't delete article, code: ", w.Code)
	}
}

func TestManageArticles(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	id := createArticle()
	//Prepare context
	prepareContextForManage(c)
	//Checking that we can get all our posts
	ManageArticles(c)

	services.DeleteArticle(id, username)

	if w.Code != http.StatusOK {
		t.Error("Error code: ", w.Code)
	}
	//Checking that we can't get our posts, if we haven't anyone
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	prepareContextForManage(c)
	//Checking that we can't get our posts, if we don't authorize
	ManageArticles(c)

	if w.Code != http.StatusNotFound {
		t.Error("Error code: ", w.Code)
	}

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/article/manage/", nil)
	c.Set("is_logged_in", false)
	c.Request.Header.Set("Accept", "application/json")

	ManageArticles(c)

	if w.Code != http.StatusUnauthorized {
		t.Error("Error code: ", w.Code)
	}

}

func TestGetUpdateArticle(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	//Create article which we want update
	id := createArticle()
	//Prepare our context
	c.Request, _ = http.NewRequest("GET", "/article/update/", nil)
	c.Set("is_logged_in", true)
	params := make([]gin.Param, 10)
	c.Params = params
	c.Params[0] = gin.Param{Key: "id", Value: id}
	c.Request.Header.Set("Accept", "application/json")
	token := GenerateToken(username, time.Duration(time.Second*3))
	c.Request.Header.Set("Cookie", "token="+token)

	GetUpdateArticle(c)

	if w.Code != http.StatusOK {
		t.Error("Status: ", w.Code, " Body: ", w.Body)
	}

	services.DeleteArticle(id, username)
}

func TestPostUpdateArticle(t *testing.T) {

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	//Create article which we want update
	id := createArticle()
	defer services.DeleteArticle(id, username)
	//Prepare our context
	c.Request, _ = http.NewRequest("POST", "/article/update/", nil)
	prepareContextForUpdate(c)
	c.Params[0] = gin.Param{Key: "id", Value: id}
	token := GenerateToken(username, time.Duration(time.Second*3))
	c.Request.Header.Set("Cookie", "token="+token)

	PostUpdateArticle(c)

	if w.Code != http.StatusOK {
		t.Error("Status: ", w.Code, " Body: ", w.Body)
	}
	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/article/update/", nil)
	prepareContextForUpdate(c)
	token = GenerateToken("1234567", time.Duration(time.Second*3))
	c.Request.Header.Set("Cookie", "token="+token)

	PostUpdateArticle(c)

	if w.Code != http.StatusBadRequest {
		t.Error("Status: ", w.Code, " Body: ", w.Body)
	}
}

func createArticle() string {
	services.SaveArticle(title, content, username)

	article := services.GetUserArticles(username)[0]
	id := fmt.Sprint(article.Id)

	return id
}

func prepareContextForManage(c *gin.Context) {
	c.Request, _ = http.NewRequest("GET", "/article/manage/", nil)
	c.Set("is_logged_in", true)
	c.Request.Header.Set("Accept", "application/json")
	token := GenerateToken(username, time.Duration(time.Second*5))
	c.Request.Header.Set("Cookie", "token="+token)
}

func prepareContext(c *gin.Context) {
	c.Request, _ = http.NewRequest("POST", "/article/create/", nil)

	c.Set("is_logged_in", true)
	c.Request.Header.Set("Accept", "application/json")
	c.Request.PostForm = url.Values{}
	c.Request.PostForm.Set("title", "title")
	c.Request.PostForm.Set("content", "content")
	token := GenerateToken("testuser", time.Duration(time.Second*5))
	c.Request.Header.Set("Cookie", "token="+token)
}

func prepareContextForUpdate(c *gin.Context) {
	c.Set("is_logged_in", true)
	params := make([]gin.Param, 10)
	c.Params = params
	c.Request.Header.Set("Accept", "application/json")
	c.Request.PostForm = url.Values{}
	c.Request.PostForm.Set("title", "title")
	c.Request.PostForm.Set("content", "content")
}
