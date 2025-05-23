package db

import (
	"github.com/teakingwang/gin-demo/internal/model"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}
}
