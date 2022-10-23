package middleware

import (
	"net/http"

	"github.com/S3B4SZ17/Email_service/services"
	"github.com/gin-gonic/gin"
)

func Oauth2AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := services.ValidateToken(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
