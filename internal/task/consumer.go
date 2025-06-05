package task

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/teakingwang/gin-demo/config"
	"github.com/teakingwang/gin-demo/pkg/mq"
)

func RegisterConsumers(mq *mq.RocketMQ) error {
	err := mq.RegisterConsumer(config.Config.RocketMQ.ConsumerTopic, func(ctx context.Context, msgs []*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range msgs {
			fmt.Printf("received message: %s\n", string(msg.Body))
		}
		return consumer.ConsumeSuccess, nil
	})

	if err != nil {
		return fmt.Errorf("register consumer failed: %w", err)
	}

	err = mq.StartConsumer()
	if err != nil {
		return fmt.Errorf("start consumer failed: %w", err)
	}

	return nil
}
