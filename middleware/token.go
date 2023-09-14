package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(username string) string {
	// Define a secret key for signing the token
	secretKey := []byte("AK8}<|Vw4>F&y!I8.O>&B}F(gd4N[i")

	// Create a new token with claims
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username

	// Set an expiration time for the token (e.g., 1 day)
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Sign the token with the secret key
	tokenString, _ := token.SignedString(secretKey)

	return tokenString
}
