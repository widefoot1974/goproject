package logger

import (
	"log"
	"os"
)

// Log levels
const (
	DebugLevel = iota
	InfoLevel
	WarningLevel
	ErrorLevel
)

type Logger struct {
	Level       int
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

var logger *Logger

// init function
func init() {
	logger = &Logger{
		Level:       InfoLevel, // just the default value
		debugLogger: log.New(os.Stdout, "DEBUG: ", log.LstdFlags|log.Lshortfile),
		infoLogger:  log.New(os.Stdout, "INFO: ", log.LstdFlags),
		warnLogger:  log.New(os.Stdout, "WARN: ", log.LstdFlags),
		errorLogger: log.New(os.Stdout, "ERROR: ", log.LstdFlags|log.Lshortfile),
	}
}

// Set log level
func SetLevel(level int) {
	logger.Level = level
}

// Methods to log (at different levels)
func Debug(message string) {
	if logger.Level <= DebugLevel {
		logger.debugLogger.Println(message)
	}
}

func Info(message string) {
	if logger.Level <= InfoLevel {
		logger.infoLogger.Println(message)
	}
}

func Warning(message string) {
	if logger.Level <= WarningLevel {
		logger.warnLogger.Println(message)
	}
}

func Error(message string) {
	if logger.Level <= ErrorLevel {
		logger.errorLogger.Println(message)
	}
}
