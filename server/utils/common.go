package utils

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func GetLog(c *gin.Context) *zap.Logger {
	logger := c.MustGet("logger").(*zap.Logger)
	return logger
}

func Log(c *gin.Context) func(msg string, fields ...zapcore.Field) {
	logger := c.MustGet("logger").(*zap.Logger)
	return logger.Info
}
