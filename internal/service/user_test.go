package service

import (
	"context"
	"github.com/teakingwang/gin-demo/config"
	"github.com/teakingwang/gin-demo/internal/repository"
	"github.com/teakingwang/gin-demo/pkg/db"
	"github.com/teakingwang/gin-demo/pkg/idgen"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// load config
	os.Setenv("MY_APP_CONFIG_PATH", "/Users/teaking/code/go/src1/gin-demo/resources")
	config.LoadConfig()
	config.Config.Database.Host = "127.0.0.1"

	// 初始化数据库连接
	gdb, err := db.NewDB()
	if err != nil {
		panic(err)
	}
	gdb.Migrator()

	err = idgen.Init()
	if err != nil {
		panic(err)
	}

	// 运行测试
	m.Run()
}

func TestUserService_CreateUser(t *testing.T) {
	gormDB, err := db.NewDB()
	if err != nil {
		panic(err)
	}

	srv := &UserService{
		userRepo: repository.NewUserRepo(gormDB),
	}
	userID, err := srv.CreateUser(context.Background(), &CreateUser{
		Mobile: "12345678901",
	})
	t.Log(userID, err)
}
