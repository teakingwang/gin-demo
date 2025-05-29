package mq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/producer"

	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

type MQProducer struct {
	producer rocketmq.Producer
	topic    string
}

func NewProducer(nameServer, topic string, group string) (*MQProducer, error) {
	p, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{nameServer}),
		producer.WithGroupName(group),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}

	if err := p.Start(); err != nil {
		return nil, fmt.Errorf("failed to start producer: %w", err)
	}

	return &MQProducer{
		producer: p,
		topic:    topic,
	}, nil
}

func (p *MQProducer) Send(ctx context.Context, msg interface{}) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("marshal message failed: %w", err)
	}
	message := primitive.NewMessage(p.topic, body)
	_, err = p.producer.SendSync(ctx, message)
	if err != nil {
		return fmt.Errorf("send message failed: %w", err)
	}
	return nil
}

func (p *MQProducer) Shutdown() error {
	return p.producer.Shutdown()
}
