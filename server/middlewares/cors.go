package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CORSHandler adds the necessary CORS headers to an answer, and handles
// the preflight OPTIONS request
func HandleCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		if origin := c.GetHeader("Origin"); origin != "" {
			c.Header("Access-Control-Allow-Origin",
				origin)
			c.Header("Access-Control-Allow-Methods",
				"POST, GET, OPTIONS, PUT, DELETE")
			c.Header("Access-Control-Allow-Headers",
				"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		}

		// Stop here if its Preflighted OPTIONS request
		if c.Request.Method == "OPTIONS" {
			c.Status(http.StatusOK)
			return
		}

		c.Next()
	}
}
