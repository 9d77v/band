package mq

import (
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
)

type RocketMQ struct {
	Conf     Conf
	Producer rocketmq.Producer
}

func NewRocketMQ(conf Conf) RocketMQ {
	p, _ := rocketmq.NewProducer(
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{conf.NamesrvAddr})),
		producer.WithRetry(2),
		producer.WithGroupName(conf.ProducerGroup),
	)
	err := p.Start()
	if err != nil {
		panic(err)
	}
	return RocketMQ{
		Conf:     conf,
		Producer: p,
	}
}
