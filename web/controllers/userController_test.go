package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/maoltr/blog/services"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestShowLoginPage(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("GET", "/u/login/", nil)

	c.Set("is_logged_in", false)
	c.Request.Header.Set("Accept", "application/json")

	ShowLoginPage(c)

	if w.Code != http.StatusOK {
		t.Error("Status: ", w.Code, " Body: ", w.Body)
	}
}

func TestShowRegisterPage(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("GET", "/u/register/", nil)

	c.Set("is_logged_in", false)
	c.Request.Header.Set("Accept", "application/json")

	ShowRegisterPage(c)

	if w.Code != http.StatusOK {
		t.Error("Status: ", w.Code, " Body: ", w.Body)
	}
}

func TestLogOut(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest("GET", "/u/logout/", nil)

	c.Set("is_logged_in", true)
	c.Request.Header.Set("Accept", "application/json")

	LogOut(c)

	if w.Code != http.StatusOK {
		t.Error("Status: ", w.Code, " Body: ", w.Body)
	}
}

func TestLogin(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	prepareContextForLogin(c)

	Login(c)

	if w.Code != http.StatusOK {
		t.Error("Status: ", w.Code, " Body: ", w.Body)
	}

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	prepareContextForLogin(c)

	c.Request.Header.Set("Accept", "application/json")
	c.Request.PostForm.Set("username", "somenonexistname,ihope")
	c.Request.PostForm.Set("password", "1")

	Login(c)

	if w.Code != http.StatusBadRequest {
		t.Error("Status: ", w.Code, " Body: ", w.Body)
	}
}

func TestRegister(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	prepareContextForRegister(c)

	Register(c)

	if w.Code != http.StatusOK {
		t.Error("Status: ", w.Code, " Body: ", w.Body)
	}

	w = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(w)
	prepareContextForRegister(c)

	c.Request.Header.Set("Accept", "application/json")
	c.Request.PostForm.Set("username", username)
	c.Request.PostForm.Set("password", "1")

	Register(c)

	if w.Code != http.StatusBadRequest {
		t.Error("Status: ", w.Code, " Body: ", w.Body)
	}

	services.DeleteUser(username)
}

func prepareContextForLogin(c *gin.Context) {
	c.Request, _ = http.NewRequest("POST", "/u/login/", nil)

	c.Set("is_logged_in", false)
	c.Request.Header.Set("Accept", "application/json")
	c.Request.PostForm = url.Values{}
	c.Request.PostForm.Set("username", "Joni")
	c.Request.PostForm.Set("password", "w3edr509bc")
}

func prepareContextForRegister(c *gin.Context) {
	c.Request, _ = http.NewRequest("POST", "/u/register/", nil)

	c.Set("is_logged_in", false)
	c.Request.Header.Set("Accept", "application/json")
	c.Request.PostForm = url.Values{}
	c.Request.PostForm.Set("username", username)
	c.Request.PostForm.Set("password", "12345678")
}
