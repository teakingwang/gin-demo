package service

import (
	"context"
	"github.com/teakingwang/gin-demo/config"
	"github.com/teakingwang/gin-demo/internal/consts"
	"github.com/teakingwang/gin-demo/internal/model"
	"github.com/teakingwang/gin-demo/internal/repository"
	"github.com/teakingwang/gin-demo/pkg/datastore/redis"
	"github.com/teakingwang/gin-demo/pkg/generator"
	"github.com/teakingwang/gin-demo/pkg/idgen"
	"go.uber.org/zap"
	"time"
)

type UserService struct {
	logger   *zap.SugaredLogger
	redis    redis.Store
	userRepo *repository.UserRepo
}

func NewUserService(redisStore redis.Store, logger *zap.SugaredLogger, userRepo *repository.UserRepo) *UserService {
	return &UserService{userRepo: userRepo, logger: logger, redis: redisStore}
}

func (s *UserService) GetAllUsers() ([]model.User, error) {
	s.logger.Info("call GetAllUsers")
	return s.userRepo.GetAllUsers()
}

func (s *UserService) CreateUser(ctx context.Context, create *CreateUser) (int64, error) {
	s.logger.Info("call CreateUser")
	err := s.checkVerifyCode(create.VerifyCode, create.Mobile)
	if err != nil {
		return 0, err
	}

	userID := idgen.NewID()
	userItem := &model.User{
		UserID:   userID,
		Username: create.Mobile,
		Nickname: generator.GenerateNickname("cn"),
		Mobile:   create.Mobile,
	}
	if err := s.userRepo.CreateUser(ctx, userItem); err != nil {
		return userID, err
	}
	return userID, nil
}

func (s *UserService) checkVerifyCode(verifyCode, mobile string) error {
	redisCode, err := s.redis.Get(consts.KeyPrefixVerifyCode + mobile)
	if err != nil {
		return err
	}

	if verifyCode != redisCode {
		return err
	}

	err = s.redis.Del(consts.KeyPrefixVerifyCode + mobile)
	if err != nil {
		s.logger.Error("del verify code err:", err)
	}
	return nil
}

func (s *UserService) SendSms(ctx context.Context, mobile string) (string, error) {
	code, err := generator.GenerateVerifyCode(6)
	if err != nil {
		return "", err
	}

	// 验证码写入redis
	err = s.redis.Set(consts.KeyPrefixVerifyCode+mobile, code, time.Duration(config.Config.SMS.CodeExpireSeconds)*time.Second)
	if err != nil {
		return "", err
	}

	s.logger.Info("send sms code to mobile:", mobile, " code:", code)
	return code, nil
}
