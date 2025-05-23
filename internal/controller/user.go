package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/teakingwang/gin-demo/internal/app"
	"github.com/teakingwang/gin-demo/internal/service"
	"github.com/teakingwang/gin-demo/pkg/errs"
	"github.com/teakingwang/gin-demo/pkg/resp"
)

type userController struct {
	srv *service.UserService
}

func NewUserController(ctx *app.AppContext) *userController {
	return &userController{
		srv: ctx.UserService,
	}
}

func (u *userController) GetUserList(c *gin.Context) {
	users, err := u.srv.GetAllUsers()
	if err != nil {
		resp.WriteError(c, errs.New(errs.CodeInvalidArgs, err.Error()))
		return
	}
	resp.WriteSuccess(c, users)
}

func (u *userController) RegisterUser(c *gin.Context) {
	req := &CreateUser{}
	if err := c.ShouldBindJSON(req); err != nil {
		resp.WriteError(c, errs.New(errs.CodeInvalidArgs, err.Error()))
		return
	}

	create := &service.CreateUser{
		Mobile:     req.Mobile,
		VerifyCode: req.VerifyCode,
	}
	id, err := u.srv.CreateUser(c, create)
	if err != nil {
		resp.WriteError(c, errs.New(errs.CodeServerError, err.Error()))
		return
	}
	resp.WriteSuccess(c, id)
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
	resp.WriteSuccess(c, code)
}
