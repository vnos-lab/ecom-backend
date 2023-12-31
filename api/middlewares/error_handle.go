package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (e *GinMiddleware) ErrorHandler(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			if c.Errors.Last().Type != gin.ErrorTypePublic {
				logger.WithOptions(zap.AddStacktrace(zap.DPanicLevel)).Error(fmt.Sprintf("Error when handling request: %+v", c.Errors.Last().Err))
			} else {
				logger.WithOptions(zap.AddStacktrace(zap.DPanicLevel)).Warn(fmt.Sprintf("Error when handling request: %+v", c.Errors.Last().Err))
			}
		}
	}
}
