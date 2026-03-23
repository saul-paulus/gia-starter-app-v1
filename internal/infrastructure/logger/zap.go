package logger

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

// InitLogger initializes the Zap logger
func InitLogger() {
	// Ensure logs directory exists
	logDir := "logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		_ = os.Mkdir(logDir, 0755)
	}

	logFile := filepath.Join(logDir, "app.log")

	// Encoder config
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// File core
	file, _ := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	fileWriter := zapcore.AddSync(file)
	fileCore := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), fileWriter, zap.InfoLevel)

	// Console core
	consoleWriter := zapcore.AddSync(os.Stdout)
	consoleCore := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), consoleWriter, zap.DebugLevel)

	// Combine cores
	core := zapcore.NewTee(fileCore, consoleCore)

	// Create logger
	Log = zap.New(core, zap.AddCaller())

	defer Log.Sync()
}

// Info logs an info message
func Info(message string, fields ...zap.Field) {
	Log.Info(message, fields...)
}

// Debug logs a debug message
func Debug(message string, fields ...zap.Field) {
	Log.Debug(message, fields...)
}

// Error logs an error message
func Error(message string, fields ...zap.Field) {
	Log.Error(message, fields...)
}

// Fatal logs a fatal message and exits
func Fatal(message string, fields ...zap.Field) {
	Log.Fatal(message, fields...)
}
