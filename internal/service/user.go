package service

import (
	"context"
	"encoding/json"
	"github.com/teakingwang/gin-demo/config"
	"github.com/teakingwang/gin-demo/internal/consts"
	"github.com/teakingwang/gin-demo/internal/model"
	"github.com/teakingwang/gin-demo/internal/repository"
	"github.com/teakingwang/gin-demo/pkg/auth"
	"github.com/teakingwang/gin-demo/pkg/crypto"
	"github.com/teakingwang/gin-demo/pkg/datastore/redis"
	"github.com/teakingwang/gin-demo/pkg/generator"
	"github.com/teakingwang/gin-demo/pkg/idgen"
	"github.com/teakingwang/gin-demo/pkg/mq"
	"go.uber.org/zap"
	"time"
)

type UserService interface {
	SetPwd(ctx context.Context, userID int64, pwd string) error
	CreateUser(ctx context.Context, create *CreateUser) (string, *UserItem, error)
	SendSms(ctx context.Context, mobile string) (string, error)
	DoCleanupTask(ctx context.Context) error
}

type userService struct {
	logger   *zap.SugaredLogger
	redis    redis.Store
	mq       *mq.RocketMQ
	userRepo repository.UserRepo
}

func NewUserService(redisStore redis.Store, logger *zap.SugaredLogger, mq *mq.RocketMQ, userRepo repository.UserRepo) UserService {
	return &userService{logger: logger, redis: redisStore, mq: mq, userRepo: userRepo}
}

func (s *userService) SetPwd(ctx context.Context, userID int64, pwd string) error {
	s.logger.Info("call SetPwd")
	hashedPwd, err := crypto.HashPassword(pwd)
	if err != nil {
		return err
	}
	s.logger.Info("set password for userID:", userID)
	_, err = s.userRepo.SetPwd(ctx, userID, hashedPwd)
	if err != nil {
		return err
	}

	return err
}

func (s *userService) CreateUser(ctx context.Context, create *CreateUser) (string, *UserItem, error) {
	s.logger.Info("call CreateUser")
	err := s.checkVerifyCode(create.VerifyCode, create.Mobile)
	if err != nil {
		return "", nil, err
	}

	userItem, err := s.checkMobileExists(ctx, create.Mobile)
	if err != nil {
		return "", nil, err
	}

	token, err := auth.GenerateToken(userItem.UserID, time.Duration(config.Config.JWT.TTLSeconds)*time.Second)
	if err != nil {
		return "", nil, err
	}

	// 发布 MQ 消息
	msg := map[string]string{"event": "user_created"}
	body, err := json.Marshal(msg)
	if err != nil {
		s.logger.Error(err)
	} else {
		_ = s.mq.SendMessage(config.Config.RocketMQ.ProducerTopic, body)
	}

	return token, userItem, nil
}

func (s *userService) checkVerifyCode(verifyCode, mobile string) error {
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

// checkMobile 检查手机号是否注册
func (s *userService) checkMobileExists(ctx context.Context, mobile string) (*UserItem, error) {
	u, err := s.userRepo.GetByMobile(ctx, mobile)
	if err != nil {
		s.logger.Error("checkMobileExists error:", err)
		return nil, err
	}

	if u != nil {
		s.logger.Info("mobile exists, userID:", u.UserID)
		return &UserItem{UserID: u.UserID, HasPwd: u.Password == ""}, nil
	}

	userID := idgen.NewID()
	userItem := &model.User{
		UserID:   userID,
		Username: mobile,
		Nickname: generator.GenerateNickname("cn"),
		Mobile:   mobile,
	}

	return &UserItem{UserID: userID, HasPwd: false}, s.userRepo.CreateUser(ctx, userItem)
}

func (s *userService) SendSms(ctx context.Context, mobile string) (string, error) {
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

func (s *userService) DoCleanupTask(ctx context.Context) error {
	s.logger.Info("cron test...")
	return nil
}
