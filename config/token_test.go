package config

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

func TestValidateToken(t *testing.T) {
	username := "Joni"
	lifeTime := time.Duration(time.Second * 4)

	token := GenerateToken(username, lifeTime)

	res, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		t.Error(err)
	}

	if !checkToken(res, token) {
		t.Error("Validation faild")
	}

	time.Sleep(time.Second * 5)

	res, err = jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err == nil {
		t.Error(err)
	}

	if !checkToken(res, token) {
		t.Error("Validation faild")
	}
}

func TestGetUsername(t *testing.T) {
	username := "Joni"
	lifeTime := time.Duration(time.Second * 3)

	token1 := GenerateToken(username, lifeTime)
	token2 := GenerateToken("User", lifeTime)

	tokUsername, err := GetUsername(token1)

	if err != nil {
		t.Error(err)
	}

	if tokUsername != username {
		t.Error("Invalid username by GetUsername")
	}

	tokUsername, err = GetUsername(token2)

	if tokUsername != "User" {
		t.Error("Invalid username by GetUsername")
	}

	time.Sleep(time.Second * 4)

	tokUsername, err = GetUsername(token1)

	if tokUsername != "" || err == nil {
		t.Error("Got username after token is expired")
	}
}

func checkToken(token *jwt.Token, tok string) bool {
	if token.Valid {
		if ValidateToken(tok) != nil {
			return false
		}

		fmt.Println("Valid token")
		return true
	}

	if ValidateToken(tok) == nil {
		return false
	}

	fmt.Println("Invalid token")
	return true
}
