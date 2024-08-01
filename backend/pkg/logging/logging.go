package logging

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logging struct {
	Logger *zap.Logger
}

func CustomCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	baseDir := "D:/FullStack/Receipe-Sharing"
	filePath := strings.TrimPrefix(caller.File, baseDir)
	enc.AppendString(fmt.Sprintf("%s:%d",filePath,caller.Line))
}

func NewLogging() *Logging {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncodeCaller = CustomCallerEncoder // Use FullCallerEncoder for full path and function names
	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	config.StacktraceKey = ""

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

	//logger := zap.New(core, zap.AddCaller())
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)) // Adjust skip value if needed
	return &Logging{
		Logger: logger,
	}
}

type LogService interface {
	ErrorLog(funcName string, err ...error)
}

func (l *Logging) ErrorLog(funcName string, err ...error) {

	l.Logger.Sugar().Errorf("%s = %v ", funcName, err)

}
