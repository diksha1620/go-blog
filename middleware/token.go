package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(k []byte, userData interface{}) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	claims := make(jwt.MapClaims)
	claims["userData"] = userData
	claims["exp"] = time.Now().Add(time.Hour * 8760).Unix()
	token.Claims = claims
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString(k)
	return tokenString, err
}

func ValidateToken(t string, k string) (*jwt.Token, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return []byte(k), nil
	})

	return token, err
}
