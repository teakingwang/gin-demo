package app

import (
	"github.com/teakingwang/gin-demo/pkg/datastore/redis"
	"github.com/teakingwang/gin-demo/pkg/db"
	"github.com/teakingwang/gin-demo/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AppContext struct {
	Redis  redis.Store
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewAppContext() *AppContext {
	gormDB, err := db.NewDB()
	if err != nil {
		panic(err)
	}

	logger.InitProductionLogger()
	defer logger.Logger.Sync() // 确保日志都写入

	return &AppContext{
		Redis:  redis.NewRedisClient(),
		DB:     gormDB,
		Logger: zap.L(),
	}
}
