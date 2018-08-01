package controllers

import (
	"github.com/gin-gonic/gin"
	. "github.com/maoltr/blog/token"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestUserStatus(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("GET", "/get", nil)

	token := GenerateToken("joni", time.Duration(time.Second*2))

	c.Request.Header.Set("Cookie", "token="+token)

	userStatus(c)

	loggedInInterface, _ := c.Get("is_logged_in")
	loggedIn := loggedInInterface.(bool)

	if !loggedIn {
		t.Error("Problem with userStatus")
	}

	time.Sleep(time.Second * 3)

	userStatus(c)

	loggedInInterface, _ = c.Get("is_logged_in")
	loggedIn = loggedInInterface.(bool)
	if loggedIn {
		t.Error("Problem with userStatus 2")
	}

}

func TestEnsureLoggedIn(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("GET", "/get", nil)
	c.Set("is_logged_in", false)

	loggedIn(c)

	if c.Writer.Status() != 401 {
		t.Error("Expected 401 status, got: ", c.Writer.Status())
	}

	c, _ = gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("GET", "/get", nil)
	c.Set("is_logged_in", true)

	loggedIn(c)

	if c.Writer.Status() != 200 {
		t.Error("Expected 200 status, got: ", c.Writer.Status())
	}

}

func TestEnsureNotLoggedIn(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("GET", "/get", nil)
	c.Set("is_logged_in", false)

	noLoggedIn(c)

	if c.Writer.Status() != 200 {
		t.Error("Expected 200 status, got: ", c.Writer.Status())
	}

	c, _ = gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("GET", "/get", nil)
	c.Set("is_logged_in", true)

	noLoggedIn(c)

	if c.Writer.Status() != 401 {
		t.Error("Expected 401 status, got: ", c.Writer.Status())
	}

}
