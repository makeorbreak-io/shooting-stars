package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/makeorbreak-io/shooting-stars/server/core"
	"github.com/makeorbreak-io/shooting-stars/server/models"
	"net/http"
)

// HandleAuthentication validates the authorization token
func HandleAuthentication(authService models.IAuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithError(http.StatusUnauthorized, core.ErrorNotLogged).
				SetType(gin.ErrorTypePublic)
			return
		}

		userID, err := authService.GetUserIDByToken(token)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, core.ErrorNotLogged).
				SetType(gin.ErrorTypePublic)
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
