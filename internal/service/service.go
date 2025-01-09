package service

import (
	"gorm.io/gorm"
	"sync"
)

var (
	factory *Factory
	once    sync.Once
)

type Factory struct {
	UserSrv *UserService
}

func NewServiceFactory() {
	once.Do(func() {
		factory = &Factory{
			UserSrv: NewUserService(),
		}
	})
}

func GetServiceFactory() *Factory {
	return factory
}
