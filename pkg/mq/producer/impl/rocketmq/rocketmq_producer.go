package rocketmq

import (
	"context"

	"github.com/9d77v/band/pkg/mq"
	prod "github.com/9d77v/band/pkg/mq/producer"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

type RocketMQProducer struct {
	Conf     prod.Conf
	Producer rocketmq.Producer
}

func NewRocketMQProducer(conf prod.Conf) (*RocketMQProducer, error) {
	p, _ := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver(conf.Addrs)),
		producer.WithRetry(conf.RetryTimes),
		producer.WithGroupName(conf.GroupName),
	)
	err := p.Start()
	return &RocketMQProducer{
		Conf:     conf,
		Producer: p,
	}, err
}

func (p *RocketMQProducer) SendSync(ctx context.Context, msgs ...*mq.Message) (*mq.SendResult, error) {
	sendMsgs := []*primitive.Message{}
	for _, v := range msgs {
		sendMsgs = append(sendMsgs, &primitive.Message{Topic: v.Topic, Body: v.Body})
	}
	sendRes, err := p.Producer.SendSync(context.Background(), sendMsgs...)
	return mq.NewSendResult(int(sendRes.Status), sendRes.MsgID), err
}

func (p *RocketMQProducer) SendAsync(ctx context.Context, callback func(ctx context.Context, result *mq.SendResult, err error), msgs ...*mq.Message) error {
	sendMsgs := []*primitive.Message{}
	for _, v := range msgs {
		sendMsgs = append(sendMsgs, &primitive.Message{Topic: v.Topic, Body: v.Body})
	}
	err := p.Producer.SendAsync(context.Background(), func(ctx context.Context, result *primitive.SendResult, err error) {
		callback(ctx, mq.NewSendResult(int(result.Status), result.MsgID), err)
	}, sendMsgs...)
	return err
}
