package utils

import (
	"errors"
	"syscall"

	"go.uber.org/zap"
)

func InitLogger() (*zap.Logger, func(logger *zap.Logger)) {
	log, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	log = log.WithOptions(zap.AddStacktrace(zap.PanicLevel))

	return log, func(log *zap.Logger) {
		err := log.Sync()
		if err != nil && !errors.Is(err, syscall.ENOTTY) {
			panic(err)
		}
	}
}
