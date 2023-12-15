package consumer_factory

import (
	"log"
	"sync"

	"github.com/9d77v/band/pkg/mq/consumer"
	"github.com/9d77v/band/pkg/mq/consumer/impl/rocketmq"
)

var (
	client consumer.Consumer
	once   sync.Once
)

var (
	TypeRocketmqPush = "rocketmq_push"
)

func NewConsumer(conf consumer.Conf) (consumer.Consumer, error) {
	var err error
	var client consumer.Consumer
	switch conf.Type {
	default:
		client, err = rocketmq.NewRocketMQPushConsumer(conf)
	}
	log.Println("consumer connected to mq:", client)
	return client, err
}

func ConsumerSingleton(conf consumer.Conf) (consumer.Consumer, error) {
	var err error
	once.Do(func() {
		client, err = NewConsumer(conf)
	})
	return client, err
}
