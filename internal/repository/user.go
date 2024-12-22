package repository

import (
	"github.com/teakingwang/gin-demo/internal/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (repo *UserRepo) Migrate() error {
	return repo.db.AutoMigrate(&models.User{})
}

func (repo *UserRepo) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := repo.db.Find(&users)
	return users, result.Error
}
