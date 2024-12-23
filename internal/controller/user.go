package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/teakingwang/gin-demo/internal/service"
	"gorm.io/gorm"
	"net/http"
)

type userController struct {
	srv *service.UserService
}

func NewUserController(db *gorm.DB) *userController {
	service.NewServiceFactory(db)
	factory := service.GetServiceFactory()
	return &userController{
		srv: factory.UserSrv,
	}
}

func (u *userController) GetUserList(c *gin.Context) {
	users, err := u.srv.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
