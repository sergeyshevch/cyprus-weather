package utils

import (
	"go.uber.org/zap"
)

func InitLogger() *zap.Logger {
	log, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	log = log.WithOptions(zap.AddStacktrace(zap.PanicLevel))

	return log
}
