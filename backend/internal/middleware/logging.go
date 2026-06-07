package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nusa/backend/internal/logger"
	"go.uber.org/zap"
)

func Logging(log *logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		requestID, _ := c.Get("RequestID")

		log.Info("Request processed",
			zap.Any("request_id", requestID),
			zap.String("method", method),
			zap.String("path", path),
			zap.Int("status", status),
			zap.Duration("latency", latency),
		)
	}
}
