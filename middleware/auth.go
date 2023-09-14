package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verify the signing method and return the secret key
			if token.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("Invalid signing method")
			}
			return []byte("AK8}<|Vw4>F&y!I8.O>&B}F(gd4N[i"), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Token is valid, you can access claims from the token like this:
		// claims := token.Claims.(jwt.MapClaims)
		// username := claims["username"].(string)

		// You can also store claims or user data in the Gin context if needed
		// c.Set("username", username)

		c.Next()
	}
}
