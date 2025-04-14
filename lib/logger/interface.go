package logger

import (
	"go.uber.org/zap"
)

// Debug ...
func Debug(message string, fields ...zap.Field) {
	Logger.Debug(message, fields...)
}

// Info ...
func Info(message string, fields ...zap.Field) {
	Logger.Info(message, fields...)
}

// Warn ...
func Warn(message string, fields ...zap.Field) {
	Logger.Warn(message, fields...)
}

// Error ...
func Error(message string, fields ...zap.Field) {
	Logger.Error(message, fields...)
}

// Fatal ...
func Fatal(message string, fields ...zap.Field) {
	Logger.Fatal(message, fields...)
}
