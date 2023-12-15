package producer_factory

import (
	"log"
	"sync"

	"github.com/9d77v/band/pkg/mq/producer"
	"github.com/9d77v/band/pkg/mq/producer/impl/rocketmq"
)

var (
	client producer.Producer
	once   sync.Once
)

var (
	TypeRocketmq = "rocketmq"
)

func NewProducer(conf producer.Conf) (producer.Producer, error) {
	var err error
	var client producer.Producer
	switch conf.Type {
	default:
		client, err = rocketmq.NewRocketMQProducer(conf)
	}
	log.Println("producer connected to mq:", client)
	return client, err
}

func ProducerSingleton(conf producer.Conf) (producer.Producer, error) {
	var err error
	once.Do(func() {
		client, err = NewProducer(conf)
	})
	return client, err
}
