package logger

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	requestLoggerKey = "logger"
)

// RequestWithLogger adds logger to request.
// It sets the logger l as a value in the Gin context with the key "logger".
// This allows other parts of the application to access the logger using the same context.
func RequestWithLogger(c *gin.Context, l *zap.Logger) {
	c.Set(requestLoggerKey, l)
}
