package repository

import (
	"github.com/teakingwang/gin-demo/internal/models"
	"github.com/teakingwang/gin-demo/pkg/db"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo() *UserRepo {
	return &UserRepo{db: db.GormDB}
}

func (repo *UserRepo) Migrate() error {
	return repo.db.AutoMigrate(&models.User{})
}

func (repo *UserRepo) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := repo.db.Find(&users)
	return users, result.Error
}
