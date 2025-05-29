// pkg/mq/types.go
package mq

import (
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

// 类型别名，方便业务层使用
type Message = primitive.MessageExt
type ConsumeResult = consumer.ConsumeResult
