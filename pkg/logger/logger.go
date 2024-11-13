// pkg/logger/logger.go
package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() {
	var err error
	Logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

func Error(err error) zap.Field {
	return zap.Error(err)
}
