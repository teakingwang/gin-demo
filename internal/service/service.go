package service

import (
	"github.com/teakingwang/gin-demo/internal/app"
	"sync"
)

var (
	factory *Factory
	once    sync.Once
)

type Factory struct {
	UserSrv *UserService
}

func NewServiceFactory(ctx *app.AppContext) {
	once.Do(func() {
		factory = &Factory{
			UserSrv: NewUserService(ctx),
		}
	})
}

func GetServiceFactory() *Factory {
	return factory
}
