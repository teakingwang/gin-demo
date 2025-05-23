package repository

import (
	"context"
	"github.com/teakingwang/gin-demo/internal/model"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(gormDB *gorm.DB) *UserRepo {
	return &UserRepo{db: gormDB}
}

func (repo *UserRepo) Migrate() error {
	return repo.db.AutoMigrate(&model.User{})
}

func (repo *UserRepo) GetAllUsers() ([]model.User, error) {
	var users []model.User
	result := repo.db.Find(&users)
	return users, result.Error
}

func (repo *UserRepo) CreateUser(ctx context.Context, item *model.User) error {
	err := repo.db.Create(item).Error
	return err
}
