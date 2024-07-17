package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func RequestIDMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()
		c.Set("RequestID", requestID)

		logger.Info("New request",
			zap.String("request_id", requestID),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
		)

		// Add the request ID to the response header for tracing
		c.Writer.Header().Set("X-Request-ID", requestID)

		c.Next()

		logger.Info("Request completed",
			zap.String("request_id", requestID),
			zap.Int("status", c.Writer.Status()),
		)
	}
}
