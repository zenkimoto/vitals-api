package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Issue a new JWT token with a given key, user, and duration
func Issue(key string, user string, id uint, duration time.Duration) (string, error) {
	byteKey := []byte(key)

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = jwt.NewNumericDate(time.Now().Add(duration * time.Second))
	claims["authorized"] = true
	claims["user"] = user
	claims["id"] = id

	return token.SignedString(byteKey)
}

// Parse a JWT token with a given key
func Parse(key string, tokenString string) (string, uint, error) {
	byteKey := []byte(key)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return byteKey, nil
	})

	if err != nil {
		return "", 0, err
	}

	var user string
	var id uint

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		user = fmt.Sprintf("%v", claims["user"])
		id = uint(claims["id"].(float64))
	}

	return user, id, nil
}
