package mq

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/teakingwang/gin-demo/config"
	"github.com/teakingwang/gin-demo/pkg/logger"
	"log"
)

type RocketMQ struct {
	Producer     rocketmq.Producer
	Consumer     rocketmq.PushConsumer
	consumerInit bool
}

func NewRocketMQ() (*RocketMQ, error) {
	logger.GetSugaredLogger().Info("nameserver:", config.Config.RocketMQ.NameServer)
	p, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{config.Config.RocketMQ.NameServer}),
		producer.WithGroupName(config.Config.RocketMQ.GroupName),
		producer.WithRetry(config.Config.RocketMQ.RetryTimes),
	)
	if err != nil {
		return nil, err
	}
	if err := p.Start(); err != nil {
		return nil, err
	}

	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName(config.Config.RocketMQ.GroupName),
		consumer.WithNameServer([]string{config.Config.RocketMQ.NameServer}),
	)
	if err != nil {
		return nil, err
	}

	return &RocketMQ{
		Producer: p,
		Consumer: c,
	}, nil
}

func (r *RocketMQ) Shutdown() {
	if r.Producer != nil {
		if err := r.Producer.Shutdown(); err != nil {
			log.Printf("[WARN] Producer shutdown failed: %v", err)
		}
	}

	if r.Consumer != nil && r.consumerInit { // ✅ 仅在初始化成功后才 Shutdown
		if err := r.Consumer.Shutdown(); err != nil {
			log.Printf("[WARN] Consumer shutdown failed: %v", err)
		}
	}
}
