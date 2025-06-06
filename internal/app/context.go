package app

import (
	"github.com/teakingwang/gin-demo/config"
	"github.com/teakingwang/gin-demo/internal/repository"
	"github.com/teakingwang/gin-demo/internal/service"
	"github.com/teakingwang/gin-demo/pkg/auth"
	"github.com/teakingwang/gin-demo/pkg/datastore/redis"
	"github.com/teakingwang/gin-demo/pkg/db"
	"github.com/teakingwang/gin-demo/pkg/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AppContext struct {
	Redis       redis.Store
	DB          *gorm.DB
	Logger      *zap.SugaredLogger
	UserService service.UserService
}

func NewAppContext() *AppContext {
	// set jwt secret
	auth.SetKey(config.Config.JWT.Secret)

	gormDB, err := db.NewDB()
	if err != nil {
		panic(err)
	}

	redisStore := redis.NewRedisClient()

	logger.InitProductionLogger()
	defer logger.Logger.Sync()
	zapLogger := zap.S()

	userRepo := repository.NewUserRepo(gormDB)
	if err := userRepo.Migrate(); err != nil {
		panic("failed to migrate database")
	}

	//// 初始化 MQ Producer
	//producer, err := mq.NewProducer(config.Config.RocketMQ.NameServer, config.Config.RocketMQ.ProducerTopic, config.Config.RocketMQ.ProducerGroup)
	//zapLogger.Info(config.Config.RocketMQ.NameServer, ":", config.Config.RocketMQ.ProducerTopic, ":", config.Config.RocketMQ.ProducerGroup)
	//if err != nil {
	//	panic(err)
	//}

	userSrv := service.NewUserService(redisStore, zapLogger, nil, userRepo)

	// 可选：启动 Consumer
	//task.StartConsumers(context.Background(), userSrv)

	return &AppContext{
		Redis:       redisStore,
		DB:          gormDB,
		Logger:      zapLogger,
		UserService: userSrv,
	}
}
