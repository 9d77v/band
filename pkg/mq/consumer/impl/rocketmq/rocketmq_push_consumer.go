package rocketmq

import (
	"context"

	"github.com/9d77v/band/pkg/mq"
	con "github.com/9d77v/band/pkg/mq/consumer"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

type RocketMQPushConsumer struct {
	Conf     con.Conf
	Consumer rocketmq.PushConsumer
}

func NewRocketMQPushConsumer(conf con.Conf) (*RocketMQPushConsumer, error) {
	c, err := rocketmq.NewPushConsumer(
		consumer.WithGroupName(conf.GroupName),
		consumer.WithNsResolver(primitive.NewPassthroughResolver(conf.Addrs)),
		consumer.WithConsumeFromWhere(consumer.ConsumeFromLastOffset),
		consumer.WithConsumerModel(consumer.MessageModel(conf.ConsumeModel)),
	)
	return &RocketMQPushConsumer{
		Conf:     conf,
		Consumer: c,
	}, err
}

func (c *RocketMQPushConsumer) Start() error {
	return c.Consumer.Start()
}

func (c *RocketMQPushConsumer) Shutdown() error {
	return c.Consumer.Shutdown()
}

func (c *RocketMQPushConsumer) Subscribe(topic string, callback func(ctx context.Context, msgs ...*mq.MessageExt) (int, error)) error {
	err := c.Consumer.Subscribe(topic, consumer.MessageSelector{}, func(ctx context.Context,
		msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		extMsgs := []*mq.MessageExt{}
		for _, v := range msgs {
			extMsgs = append(extMsgs, &mq.MessageExt{
				MsgID:   v.MsgId,
				Message: *mq.NewMessage(v.Topic, v.Body),
			})
		}
		res, err := callback(ctx, extMsgs...)
		return consumer.ConsumeResult(res), err
	})
	if err != nil {
		return err
	}
	return c.Consumer.Start()
}
