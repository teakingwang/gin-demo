package repository

import (
	"context"
	"github.com/teakingwang/gin-demo/internal/model"
	"gorm.io/gorm"
)

type UserRepo interface {
	Migrate() error
	GetByMobile(ctx context.Context, mobile string) (*model.User, error)
	SetPwd(ctx context.Context, userID int64, pwd string) (int64, error)
	CreateUser(ctx context.Context, item *model.User) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(gormDB *gorm.DB) UserRepo {
	return &userRepo{db: gormDB}
}

func (repo *userRepo) Migrate() error {
	return repo.db.AutoMigrate(&model.User{})
}

func (repo *userRepo) GetByMobile(ctx context.Context, mobile string) (*model.User, error) {
	u := &model.User{}
	err := repo.db.Where("mobile = ?", mobile).First(u).Error
	if gorm.ErrRecordNotFound == err {
		return nil, nil
	}
	return u, err
}

func (repo *userRepo) SetPwd(ctx context.Context, userID int64, pwd string) (int64, error) {
	result := repo.db.Model(&model.User{}).Where("user_id = ?", userID).Update("password", pwd)
	if result.Error != nil {
		return 0, nil
	}

	return result.RowsAffected, nil
}

func (repo *userRepo) CreateUser(ctx context.Context, item *model.User) error {
	err := repo.db.Create(item).Error
	return err
}
