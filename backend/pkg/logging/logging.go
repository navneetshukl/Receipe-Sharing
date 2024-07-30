package logging

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logging struct {
}

func NewLogging() *Logging {
	return &Logging{}
}

type LogService interface {
	ErrorLog(file string, err error)
}

func (l *Logging) ErrorLog(file string, err error) {

	key := fmt.Sprintf("%s %s %v ", file, " Error is := ", err)

	fileEncoderConfig := zap.NewProductionEncoderConfig()
	fileEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoderConfig.MessageKey = "msg" // Ensures only message is included
	fileEncoder := zapcore.NewJSONEncoder(fileEncoderConfig)
	consoleEncoderConfig := zap.NewDevelopmentEncoderConfig()
	consoleEncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoderConfig.MessageKey = "msg" // Ensures only message is included
	consoleEncoder := zapcore.NewConsoleEncoder(consoleEncoderConfig)
	logFile, err := os.OpenFile("./logs/text.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	writer := zapcore.AddSync(logFile)
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)

	logger := zap.New(core, zap.AddCaller())

	logger.Error(key)
}
