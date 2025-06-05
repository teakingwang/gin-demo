package app

import (
	"github.com/teakingwang/gin-demo/config"
	"github.com/teakingwang/gin-demo/internal/repository"
	"github.com/teakingwang/gin-demo/internal/service"
	"github.com/teakingwang/gin-demo/pkg/auth"
	"github.com/teakingwang/gin-demo/pkg/datastore/redis"
	"github.com/teakingwang/gin-demo/pkg/db"
	"github.com/teakingwang/gin-demo/pkg/logger"
	"github.com/teakingwang/gin-demo/pkg/mq"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AppContext struct {
	Redis       redis.Store
	DB          *gorm.DB
	Logger      *zap.SugaredLogger
	UserService service.UserService
	MQ          *mq.RocketMQ
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
	sugar := logger.GetSugaredLogger()
	defer sugar.Sync()

	userRepo := repository.NewUserRepo(gormDB)
	if err := userRepo.Migrate(); err != nil {
		panic("failed to migrate database")
	}

	//mqClient, err := mq.NewRocketMQ()
	//if err != nil {
	//    panic(err)
	//}
	//defer mqClient.Shutdown()

	userSrv := service.NewUserService(redisStore, sugar, nil, userRepo)

	return &AppContext{
		Redis:       redisStore,
		DB:          gormDB,
		Logger:      sugar,
		UserService: userSrv,
		//  MQ:          mqClient,
	}
}
