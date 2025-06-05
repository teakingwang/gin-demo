package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

var (
	logger      *zap.Logger
	sugarLogger *zap.SugaredLogger
	once        sync.Once
)

func InitProductionLogger() {
	once.Do(func() {
		config := zap.NewProductionConfig()
		config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

		var err error
		logger, err = config.Build()
		if err != nil {
			panic("failed to initialize zap logger: " + err.Error())
		}

		sugarLogger = logger.Sugar()
		zap.ReplaceGlobals(logger)
	})
}

func GetLogger() *zap.Logger {
	if logger == nil {
		return zap.NewNop()
	}
	return logger
}

func GetSugaredLogger() *zap.SugaredLogger {
	if sugarLogger == nil {
		return zap.NewNop().Sugar()
	}
	return sugarLogger
}
