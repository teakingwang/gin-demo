package service

import (
	"context"
	"github.com/teakingwang/gin-demo/config"
	"github.com/teakingwang/gin-demo/internal/app"
	"github.com/teakingwang/gin-demo/internal/model"
	"github.com/teakingwang/gin-demo/internal/repository"
	"github.com/teakingwang/gin-demo/pkg/errs"
	"github.com/teakingwang/gin-demo/pkg/generator"
	"github.com/teakingwang/gin-demo/pkg/idgen"
	"time"
)

type UserService struct {
	ctx      *app.AppContext
	userRepo *repository.UserRepo
}

func NewUserService(ctx *app.AppContext) *UserService {
	userRepo := repository.NewUserRepo(ctx.DB)
	if err := userRepo.Migrate(); err != nil {
		panic("failed to migrate database")
	}
	return &UserService{userRepo: userRepo, ctx: ctx}
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	s.ctx.Logger.Info("call GetAllUsers")
	return s.userRepo.GetAllUsers()
}

func (s *UserService) CreateUser(ctx context.Context, create *CreateUser) (int64, error) {
	s.ctx.Logger.Info("call CreateUser")
	userID := idgen.NewID()
	userItem := &model.User{
		UserID:   userID,
		Username: create.Mobile,
		Nickname: generator.GenerateNickname("cn"),
		Mobile:   create.Mobile,
	}
	if err := s.userRepo.CreateUser(ctx, userItem); err != nil {
		return userID, errs.New(errs.CodeServerError, err.Error())
	}
	return userID, nil
}

func (s *UserService) checkVerifyCode(verifyCode, mobile string) error {
	if verifyCode != "123456" {
		return errs.New(errs.CodeInvalidArgs, "verify code is invalid")
	}
	return nil
}

func (s *UserService) SendSms(ctx context.Context, mobile string) (string, error) {
	code, err := generator.GenerateVerifyCode(6)
	if err != nil {
		return "", errs.New(errs.CodeServerError, err.Error())
	}

	// 验证码写入redis
	s.ctx.Redis.Set(mobile, code, time.Duration(config.Config.SMS.CodeExpireSeconds))
	return "", nil
}
