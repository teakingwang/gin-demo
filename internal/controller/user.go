package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/teakingwang/gin-demo/internal/app"
	"github.com/teakingwang/gin-demo/internal/service"
	"github.com/teakingwang/gin-demo/pkg/errs"
	"github.com/teakingwang/gin-demo/pkg/resp"
	"net/http"
)

type userController struct {
	srv *service.UserService
}

func NewUserController(ctx *app.AppContext) *userController {
	service.NewServiceFactory(ctx)
	factory := service.GetServiceFactory()
	return &userController{
		srv: factory.UserSrv,
	}
}

func (u *userController) GetUserList(c *gin.Context) {
	users, err := u.srv.GetAllUsers()
	if err != nil {
		resp.WriteError(c, errs.New(errs.CodeInvalidArgs, err.Error()))
		return
	}
	c.JSON(http.StatusOK, users)
}

func (u *userController) RegisterUser(c *gin.Context) {
	req := &CreateUser{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp.WriteError(c, errs.New(errs.CodeInvalidArgs, err.Error()))
		return
	}

	create := &service.CreateUser{
		Mobile: req.Mobile,
	}
	id, err := u.srv.CreateUser(c, create)
	if err != nil {
		resp.WriteError(c, errs.New(errs.CodeServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, id)
}

func (u *userController) SendSms(c *gin.Context) {
	req := &SendSms{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp.WriteError(c, errs.New(errs.CodeInvalidArgs, err.Error()))
		return
	}

	code, err := u.srv.SendSms(c, req.Mobile)
	if err != nil {
		resp.WriteError(c, errs.New(errs.CodeServerError, err.Error()))
		return
	}
	c.JSON(http.StatusOK, code)
}
