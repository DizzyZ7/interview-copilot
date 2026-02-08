package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Middleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		if !strings.HasPrefix(h, "Bearer ") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		userID, err := Parse(strings.TrimPrefix(h, "Bearer "), secret)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("user_id", userID)
		c.Next()
	}
}
