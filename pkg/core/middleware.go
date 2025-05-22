package core

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			statusCode := http.StatusInternalServerError
			if c.Writer.Status() != http.StatusOK {
				statusCode = c.Writer.Status()
			}

			if statusCode == http.StatusInternalServerError {
				env := os.Getenv("ENVIRONMENT")
				if env == "development" {
					c.Error(c.Errors[0])
					c.Abort()
				} else {
					NewErrorResponse(c, http.StatusInternalServerError, "Internal Server Error", nil)
				}
			}
		}
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, NewErrorResponse(c, http.StatusUnauthorized, "Unauthorized", "Missing token"))
			return
		}

		token := strings.Replace(authHeader, "Bearer ", "", 1)

		if token != "valid" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, NewErrorResponse(c, http.StatusUnauthorized, "Unauthorized", "Invalid token"))
			return
		}

		c.Next()
	}
}
