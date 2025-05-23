// pkg/logger/logger.go
package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitProductionLogger() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

	var err error
	Logger, err = config.Build()
	if err != nil {
		panic("failed to initialize zap logger: " + err.Error())
	}

	zap.ReplaceGlobals(Logger) // 设置 zap.L() 默认使用这个 logger
}
