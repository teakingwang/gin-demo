package service

import (
	"github.com/teakingwang/gin-demo/internal/models"
	"github.com/teakingwang/gin-demo/internal/repository"
)

type UserService struct {
	userRepo *repository.UserRepo
}

func NewUserService() *UserService {
	userRepo := repository.NewUserRepo()
	if err := userRepo.Migrate(); err != nil {
		panic("failed to migrate database")
	}
	return &UserService{userRepo: userRepo}
}

func (service *UserService) GetAllUsers() ([]models.User, error) {
	return service.userRepo.GetAllUsers()
}
