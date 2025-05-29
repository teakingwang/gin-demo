package service

import (
	"context"
	"encoding/json"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"go.uber.org/zap"
)

type UserRegisterMessage struct {
	UserID uint   `json:"user_id"`
	Phone  string `json:"phone"`
	// 其他字段
}

func (s *userService) HandleUserMessage(ctx context.Context, msg *primitive.MessageExt) consumer.ConsumeResult {
	s.logger.Info("received user message", zap.ByteString("body", msg.Body))

	var data UserRegisterMessage
	if err := json.Unmarshal(msg.Body, &data); err != nil {
		s.logger.Error("failed to unmarshal message", zap.Error(err))
		return consumer.ConsumeRetryLater
	}

	// 处理业务逻辑
	s.logger.Info("handling user message", zap.Uint("user_id", data.UserID))

	// 示例：写日志、发通知、调用其他 service 等

	return consumer.ConsumeSuccess
}
