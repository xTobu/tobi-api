package middleware

import (
	"github.com/gin-gonic/gin"

	"tobi-api/lib/log"
)

// Auth : Auth Middleware
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Warn("AbortWithStatus(500)")
		c.AbortWithStatus(500)
		c.Next()
	}
}
