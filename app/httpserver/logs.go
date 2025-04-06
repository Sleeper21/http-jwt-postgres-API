package httpserver

import (
	"core/app/domain"
	"time"

	"strings"

	"github.com/gin-gonic/gin"
)

func loggerMiddleware(logger domain.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now().UTC()
		// Process request
		c.Next()
		elapsed := time.Since(start)
		sanitizedURL := strings.ReplaceAll(c.Request.URL.String(), "\n", "")
		sanitizedURL = strings.ReplaceAll(sanitizedURL, "\r", "")
		logger.Debugf("Request: %s %s, Response time: %v, Status: %d", c.Request.Method,
			sanitizedURL, elapsed, c.Writer.Status())
	}
}
