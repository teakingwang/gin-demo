package task

import (
	"context"
	"github.com/teakingwang/gin-demo/config"
	"github.com/teakingwang/gin-demo/internal/service"
	"github.com/teakingwang/gin-demo/pkg/mq"
	"log"
)

func StartConsumers(ctx context.Context, userService service.UserService) {
	_, err := mq.NewConsumer(
		config.Config.RocketMQ.NameServer,
		config.Config.RocketMQ.ConsumerTopic,
		config.Config.RocketMQ.ConsumerGroup,
		userService.HandleUserMessage,
	)
	if err != nil {
		log.Fatalf("failed to start consumer: %v", err)
	}
}
