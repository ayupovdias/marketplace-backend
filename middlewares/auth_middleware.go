package middlewares

import (
	"github.com/gin-gonic/gin"
	"marketplace/utils"
	"net/http"
)

var jwtKey = []byte("secret")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := utils.ExtractToken(c.GetHeader("Authorization"))

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		token, claims, err := utils.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("user_id", uint(claims["user_id"].(float64)))
		c.Set("username", claims["username"])

		c.Next()
	}
}
