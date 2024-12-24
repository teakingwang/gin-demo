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

func NewServiceFactory(db *gorm.DB) {
	once.Do(func() {
		factory = &Factory{
			UserSrv: NewUserService(db),
		}
	})
}

func GetServiceFactory() *Factory {
	return factory
}
