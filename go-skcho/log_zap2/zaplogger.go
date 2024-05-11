package zaplogger

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	colorRed     = "\033[31m"
	colorGreen   = "\033[32m"
	colorYellow  = "\033[33m"
	colorBlue    = "\033[34m"
	colorMagenta = "\033[35m"
	colorReset   = "\033[0m"
)

var (
	logger       *zap.Logger
	consoleLevel zap.AtomicLevel
	fileLevel    zap.AtomicLevel
	logDirPath   string = "/home/enfr/log" // Default value
	logFileName  string = "zap_log"        // Default value
)

type DailyRotatingFile struct {
	currentDate string
	file        *os.File
}

func SetLogDirPath(path string) {
	logDirPath = path
}

func SetLogFileName(name string) {
	logFileName = name
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("[2006-01-02 15:04:05.000]"))
}

func customLevelEncoderConsole(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	var levelColor string
	switch level {
	case zapcore.DebugLevel:
		levelColor = colorBlue
	case zapcore.InfoLevel:
		levelColor = colorGreen
	case zapcore.WarnLevel:
		levelColor = colorYellow
	case zapcore.ErrorLevel:
		levelColor = colorRed
	case zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		levelColor = colorMagenta
	default:
		levelColor = colorReset
	}
	enc.AppendString(fmt.Sprintf("%s%-5s%s", levelColor, strings.ToUpper(level.String()), colorReset))
}

func customLevelEncoderFile(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("%-5s", strings.ToUpper(level.String())))
}

func customCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("%s:%d", filepath.Base(caller.File), caller.Line))
}

func NewDailyRotatingFile() *DailyRotatingFile {
	initialDate := time.Now().Format("2006-01-02")
	logFile := fmt.Sprintf("%s.%s.log", logFileName, initialDate)
	logFilePath := filepath.Join(logDirPath, logFile)
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("os.OpenFile(%v) fail: %v\n", logFilePath, err)
		return nil
	}
	return &DailyRotatingFile{
		currentDate: initialDate,
		file:        file,
	}
}

func (drf *DailyRotatingFile) Write(p []byte) (n int, err error) {
	now := time.Now().Format("2006-01-02")
	if now != drf.currentDate {
		drf.Rotate(now)
	}
	return drf.file.Write(p)
}

func (drf *DailyRotatingFile) Sync() error {
	return drf.file.Sync()
}

func (drf *DailyRotatingFile) Rotate(newDate string) {
	drf.file.Close()
	var err error
	logFile := fmt.Sprintf("%s.%s.log", logFileName, newDate)
	logFilePath := filepath.Join(logDirPath, logFile)
	drf.file, err = os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("os.OpenFile(%v) fail: %v\n", logFilePath, err)
	}
	drf.currentDate = newDate
}

func InitZapLogger() {
	consoleEncoderConfig := zap.NewProductionEncoderConfig()
	consoleEncoderConfig.EncodeTime = customTimeEncoder
	consoleEncoderConfig.EncodeCaller = customCallerEncoder
	consoleEncoderConfig.EncodeLevel = customLevelEncoderConsole
	consoleEncoderConfig.ConsoleSeparator = " "

	fileEncoderConfig := zap.NewProductionEncoderConfig()
	fileEncoderConfig.EncodeTime = customTimeEncoder
	fileEncoderConfig.EncodeCaller = customCallerEncoder
	fileEncoderConfig.EncodeLevel = customLevelEncoderFile
	fileEncoderConfig.ConsoleSeparator = " "

	consoleLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
	fileLevel = zap.NewAtomicLevelAt(zap.DebugLevel)

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(consoleEncoderConfig), zapcore.Lock(os.Stdout), consoleLevel),
		zapcore.NewCore(zapcore.NewConsoleEncoder(fileEncoderConfig), zapcore.AddSync(NewDailyRotatingFile()), fileLevel),
	)

	logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

// func SetZapLogLevel(cLevel zapcore.Level, fLevel zapcore.Level) {
//         consoleLevel.SetLevel(cLevel)
//         fileLevel.SetLevel(fLevel)
// }

func SetLogLevel(consoleLevelStr, fileLevelStr string) {
	cLevel, err := parseLogLevel(consoleLevelStr)
	if err != nil {
		Errorf("ConsoleLogLevel is invalid (%v), then default InfoLevel", consoleLevelStr)
		cLevel = zapcore.InfoLevel
	}

	fLevel, err := parseLogLevel(fileLevelStr)
	if err != nil {
		Errorf("FileLogLevel is invalid (%v), then default DebugLevel", fileLevelStr)
		cLevel = zapcore.DebugLevel
	}

	consoleLevel.SetLevel(cLevel)
	fileLevel.SetLevel(fLevel)
	return nil
}

func parseLogLevel(levelStr string) (zapcore.Level, error) {
	switch strings.ToLower(levelStr) {
	case "debug":
		return zapcore.DebugLevel, nil
	case "info":
		return zapcore.InfoLevel, nil
	case "warn", "warning":
		return zapcore.WarnLevel, nil
	case "error":
		return zapcore.ErrorLevel, nil
	default:
		return zapcore.InfoLevel, fmt.Errorf("unknown log level: %s", levelStr)
	}
}

func Logf(level zapcore.Level, callerSkip int, template string, args ...interface{}) {
	logger.WithOptions(zap.AddCallerSkip(callerSkip)).Check(level, fmt.Sprintf(template, args...)).Write()
}

func Debugf(template string, args ...interface{}) {
	Logf(zapcore.DebugLevel, 1, template, args...)
}

func Infof(template string, args ...interface{}) {
	Logf(zapcore.InfoLevel, 1, template, args...)
}

func Warnf(template string, args ...interface{}) {
	Logf(zapcore.WarnLevel, 1, template, args...)
}

func Errorf(template string, args ...interface{}) {
	Logf(zapcore.ErrorLevel, 1, template, args...)
}

func DPanicf(template string, args ...interface{}) {
	Logf(zapcore.DPanicLevel, 1, template, args...)
}

func Panicf(template string, args ...interface{}) {
	Logf(zapcore.PanicLevel, 1, template, args...)
}

func Fatalf(template string, args ...interface{}) {
	Logf(zapcore.FatalLevel, 1, template, args...)
}
