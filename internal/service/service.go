package service

import (
	"gorm.io/gorm"
	"sync"
)

var (
	factory *serviceFactory
	once    sync.Once
)

type serviceFactory struct {
	UserSrv *UserService
}

func NewServiceFactory(db *gorm.DB) *serviceFactory {
	once.Do(func() {
		factory = &serviceFactory{
			UserSrv: NewUserService(db),
		}
	})
	return factory
}
