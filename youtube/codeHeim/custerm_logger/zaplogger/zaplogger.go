package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log levels
const (
	DebugLevel = iota
	InfoLevel
	WarningLevel
	ErrorLevel
)

var logger *zap.Logger
var atomicLevel zap.AtomicLevel

// init function to setup the logger with default level as InfoLevel
func init() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // setting time encoder for the logs

	// Set initial level to InfoLevel by default
	atomicLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
	config.Level = atomicLevel

	var err error
	logger, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
}

// SetLevel adjusts the logging level at runtime
func SetLevel(level int) {
	var zapLevel zapcore.Level
	switch level {
	case DebugLevel:
		zapLevel = zap.DebugLevel
	case InfoLevel:
		zapLevel = zap.InfoLevel
	case WarningLevel:
		zapLevel = zap.WarnLevel
	case ErrorLevel:
		zapLevel = zap.ErrorLevel
	default:
		zapLevel = zap.InfoLevel
	}
	atomicLevel.SetLevel(zapLevel)
}

// Debug logs a message at DebugLevel
func Debug(message string) {
	logger.Debug(message)
}

// Info logs a message at InfoLevel
func Info(message string) {
	logger.Info(message)
}

// Warning logs a message at WarningLevel
func Warning(message string) {
	logger.Warn(message)
}

// Error logs a message at ErrorLevel
func Error(message string) {
	logger.Error(message)
}
