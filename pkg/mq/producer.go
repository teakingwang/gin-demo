package mq

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func (r *RocketMQ) SendMessage(topic string, body []byte) error {
	msg := primitive.NewMessage(topic, body)
	_, err := r.Producer.SendSync(context.Background(), msg)
	return err
}
