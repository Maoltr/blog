package config

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"time"
)

var mySigningKey = []byte("secret")

func GenerateToken(username string, lifeTime time.Duration) string {
	var token jwt.Token

	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(lifeTime).Unix(),
	}

	token = *jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, _ := token.SignedString(mySigningKey)

	return tokenString
}

func ValidateToken(token string) error {
	res, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err == nil {
		if res.Valid {
			return nil
		}
	}

	return errors.New("Not valid token")
}

func GetUsername(token string) (string, error) {
	res, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err == nil {
		if res.Valid {
			claims := res.Claims.(jwt.MapClaims)

			return claims["username"].(string), nil

		}
	}
	fmt.Println(err.Error())
	return "", errors.New("Not valid token")
}
