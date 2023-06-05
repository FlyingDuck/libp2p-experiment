package logger

import (
	"log"

	"go.uber.org/zap"
)

var lgger *zap.Logger

func init() {
	var err error
	lgger, err = zap.NewProduction(zap.AddCallerSkip(1))
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
}

func Info(msg string, fields ...zap.Field) {
	lgger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	lgger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	lgger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	lgger.Fatal(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	lgger.Panic(msg, fields...)
}
