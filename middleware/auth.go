package middleware

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthHandlerAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		// Check if toke in correct format
		// ie Bearer: xx03xllasx
		b := "Bearer "
		if !strings.Contains(token, b) {
			c.JSON(403, gin.H{"message": "Your request is not authorized", "status": 403})
			c.Abort()
			return
		}
		t := strings.Split(token, b)
		if len(t) < 2 {
			c.JSON(403, gin.H{"message": "An authorization token was not supplied", "status": 403})
			c.Abort()
			return
		}
		// Validate token
		valid, err := ValidateToken(t[1], SigningKey)
		if err != nil {
			c.JSON(403, gin.H{"message": "Invalid authorization token", "status": 403})
			c.Abort()
			return
		}

		// set userId Variable
		c.Set("userData", valid.Claims.(jwt.MapClaims)["userData"])
		c.Next()
	}
}

func AuthHandlerAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		// Check if toke in correct format
		// ie Bearer: xx03xllasx
		b := "Bearer "
		if !strings.Contains(token, b) {
			c.JSON(403, gin.H{"message": "Your request is not authorized", "status": 403})
			c.Abort()
			return
		}
		t := strings.Split(token, b)
		if len(t) < 2 {
			c.JSON(403, gin.H{"message": "An authorization token was not supplied", "status": 403})
			c.Abort()
			return
		}
		// Validate token
		valid, err := ValidateToken(t[1], SigningKey)
		if err != nil {
			c.JSON(403, gin.H{"message": "Invalid authorization token", "status": 403})
			c.Abort()
			return
		}

		// set userId Variable
		c.Set("userData", valid.Claims.(jwt.MapClaims)["userData"])
		c.Next()
	}
}
