package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewZap(logLevel zapcore.Level) *zap.Logger {
	config := zap.NewProductionEncoderConfig()
	config.TimeKey = "T"
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncodeLevel = zapcore.CapitalColorLevelEncoder

	encoder := zapcore.NewConsoleEncoder(config)

	return zap.New(zapcore.NewCore(encoder, os.Stdout, logLevel))
}
