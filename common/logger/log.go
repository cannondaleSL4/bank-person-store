package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func Init(serviceName string) {
	var err error

	// Добавляем имя сервиса к логгеру
	initialFields := make(map[string]interface{})
	initialFields["service"] = serviceName

	productionConfig := zap.NewProductionConfig()
	productionConfig.EncoderConfig.TimeKey = "timestamp"
	productionConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	log, err = productionConfig.Build(zap.Fields(zap.String("service", serviceName)))
	if err != nil {
		panic(err)
	}
}

func GetLogger() *zap.Logger {
	return log
}

func Sync() {
	log.Sync()
}
