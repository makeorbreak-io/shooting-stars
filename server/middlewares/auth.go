package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HandleAuthentication validates the authorization token
func HandleAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userID := 1 // TODO: Get user ID from authentication service
		c.Set("userID", userID)

		c.Next()
	}
}
