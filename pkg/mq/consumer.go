package mq

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

type MessageHandler func(ctx context.Context, msgs []*primitive.MessageExt) (consumer.ConsumeResult, error)

func (r *RocketMQ) RegisterConsumer(topic string, handler MessageHandler) error {
	err := r.Consumer.Subscribe(
		topic,
		consumer.MessageSelector{},
		func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
			return handler(ctx, msgs)
		},
	)

	if err != nil {
		return err
	}

	err = r.Consumer.Start()
	if err != nil {
		return err
	}

	r.consumerInit = true // ✅ 只有成功启动后才设置为 true
	return nil
}

func (r *RocketMQ) StartConsumer() error {
	return r.Consumer.Start()
}
