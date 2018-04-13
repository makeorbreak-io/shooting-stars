package middlewares

import (
	"github.com/gin-gonic/gin"
	"log"
)

// HandleExecutionErrors checks if the request resulted in errors and displays them
func HandleExecutionErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if len(c.Errors) == 0 {
				c.Next()
				return
			}

			var errorStrings []string
			for _, err := range c.Errors {
				switch err.Type {
				case gin.ErrorTypePublic:
					errorStrings = append(errorStrings, err.Error())
				case gin.ErrorTypeRender:
					log.Printf("Error Render: %v", err)
				case gin.ErrorTypePrivate:
					log.Printf("Error: %v", err)
				}
			}

			c.JSON(c.Writer.Status(), struct {
				Errors []string `json:"errors"`
			}{
				Errors: errorStrings,
			})
		}()
		c.Next()
	}
}
