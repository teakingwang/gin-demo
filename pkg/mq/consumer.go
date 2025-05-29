package mq

import (
	"context"
	"fmt"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

type MQConsumer struct {
	consumer rocketmq.PushConsumer
}

func NewConsumer(
	nameServer, topic, group string,
	handler func(ctx context.Context, msg *primitive.MessageExt) consumer.ConsumeResult,
) (rocketmq.PushConsumer, error) {
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName(group),
		consumer.WithNameServer([]string{nameServer}),
	)
	if err != nil {
		return nil, fmt.Errorf("create consumer failed: %w", err)
	}

	err = c.Subscribe(topic, consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for _, msg := range msgs {
			result := handler(ctx, msg)
			if result != consumer.ConsumeSuccess {
				return result, nil
			}
		}
		return consumer.ConsumeSuccess, nil
	})
	if err != nil {
		return nil, fmt.Errorf("subscribe failed: %w", err)
	}

	if err := c.Start(); err != nil {
		return nil, fmt.Errorf("start consumer failed: %w", err)
	}

	return c, nil
}

func (c *MQConsumer) Shutdown() error {
	return c.consumer.Shutdown()
}
