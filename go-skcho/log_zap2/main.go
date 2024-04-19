package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	// enc.AppendString(t.Format("[2006-01-02 15:04:05.000]")) // 사용자 정의 날짜 형식
	enc.AppendString(t.Format("[01-02 15:04:05.000]")) // 사용자 정의 날짜 형식
}

type dailyRotatingFile struct {
	currentDate string
	file        *os.File
}

func (drf *dailyRotatingFile) Write(p []byte) (n int, err error) {
	now := time.Now().Format("2006-01-02")
	if now != drf.currentDate {
		drf.Rotate(now)
	}
	return drf.file.Write(p)
}

func (drf *dailyRotatingFile) Sync() error {
	return drf.file.Sync()
}

func (drf *dailyRotatingFile) Rotate(newDate string) {
	drf.file.Close() // Close the current file if it is open.
	var err error
	// drf.file, err = os.OpenFile(filepath.Join("logs", fmt.Sprintf("logfile-%s.log", newDate)), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	drf.file, err = os.OpenFile(fmt.Sprintf("logfile-%s.log", newDate), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err) // Handle error according to your needs.
	}
	drf.currentDate = newDate
}

func newDailyRotatingFile() *dailyRotatingFile {
	initialDate := time.Now().Format("2006-01-02")
	file, err := os.OpenFile(fmt.Sprintf("logfile-%s.log", initialDate), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err) // Handle error according to your needs.
	}
	return &dailyRotatingFile{
		currentDate: initialDate,
		file:        file,
	}
}

const (
	colorRed     = "\033[31m"
	colorGreen   = "\033[32m"
	colorYellow  = "\033[33m"
	colorBlue    = "\033[34m"
	colorMagenta = "\033[35m"
	colorReset   = "\033[0m"
)

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
	// enc.AppendString(fmt.Sprintf("%s%s%s", levelColor, strings.ToUpper(level.String()), colorReset))
	enc.AppendString(fmt.Sprintf("%s%-5s%s", levelColor, strings.ToUpper(level.String()), colorReset))
}

func customLevelEncoderFile(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	// enc.AppendString(strings.ToUpper(level.String()))
	enc.AppendString(fmt.Sprintf("%-5s", strings.ToUpper(level.String())))
}

func customCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	// 파일 이름만 출력하도록 설정
	enc.AppendString(fmt.Sprintf("%s:%d", filepath.Base(caller.File), caller.Line))
}

func initLogger() {
	consoleEncoderConfig := zap.NewProductionEncoderConfig()
	consoleEncoderConfig.EncodeTime = customTimeEncoder
	consoleEncoderConfig.EncodeCaller = customCallerEncoder
	consoleEncoderConfig.EncodeLevel = customLevelEncoderConsole
	consoleEncoderConfig.ConsoleSeparator = " " // 콘솔 출력 구분자 설정

	fileEncoderConfig := zap.NewProductionEncoderConfig()
	fileEncoderConfig.EncodeTime = customTimeEncoder
	fileEncoderConfig.EncodeCaller = customCallerEncoder
	fileEncoderConfig.EncodeLevel = customLevelEncoderFile
	fileEncoderConfig.ConsoleSeparator = " " // 파일 출력 구분자 설정

	core := zapcore.NewTee(
		zapcore.NewCore(zapcore.NewConsoleEncoder(consoleEncoderConfig), zapcore.Lock(os.Stdout), zapcore.DebugLevel),
		zapcore.NewCore(zapcore.NewConsoleEncoder(fileEncoderConfig), zapcore.AddSync(newDailyRotatingFile()), zapcore.DebugLevel),
	)

	logger = zap.New(core, zap.AddCaller())
}

func Logf(level zapcore.Level, template string, args ...interface{}) {
	msg := fmt.Sprintf(template, args...)
	switch level {
	case zapcore.DebugLevel:
		logger.Debug(msg)
	case zapcore.InfoLevel:
		logger.Info(msg)
	case zapcore.WarnLevel:
		logger.Warn(msg)
	case zapcore.ErrorLevel:
		logger.Error(msg)
	case zapcore.DPanicLevel:
		logger.DPanic(msg)
	case zapcore.PanicLevel:
		logger.Panic(msg)
	case zapcore.FatalLevel:
		logger.Fatal(msg)
	}
}

func Debugf(template string, args ...interface{}) {
	Logf(zapcore.DebugLevel, template, args...)
}

func Infof(template string, args ...interface{}) {
	Logf(zapcore.InfoLevel, template, args...)
}

func Warnf(template string, args ...interface{}) {
	Logf(zapcore.WarnLevel, template, args...)
}

func Errorf(template string, args ...interface{}) {
	Logf(zapcore.ErrorLevel, template, args...)
}

func DPanicf(template string, args ...interface{}) {
	Logf(zapcore.DPanicLevel, template, args...)
}

func Panicf(template string, args ...interface{}) {
	Logf(zapcore.PanicLevel, template, args...)
}

func Fatalf(template string, args ...interface{}) {
	Logf(zapcore.FatalLevel, template, args...)
}

func main() {
	initLogger()
	defer logger.Sync()

	User := struct {
		Name string
		Age  int
	}{
		Name: "John",
		Age:  51,
	}

	Infof("User %s has logged in from IP %s at %d", "Alice", "192.168.1.1", 20230101)
	Warnf("Payment of $%.2f has been processed for account %d", 123.45, 987654)
	Debugf("User = %#v", User)
	Errorf("User = %#v", User)
}
