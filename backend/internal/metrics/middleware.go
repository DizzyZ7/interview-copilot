package metrics

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		status := strconv.Itoa(c.Writer.Status())
		HttpRequests.WithLabelValues(c.Request.Method, c.FullPath(), status).Inc()
	}
}
