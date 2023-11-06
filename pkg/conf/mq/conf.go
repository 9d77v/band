package mq

import "github.com/9d77v/band/pkg/env"

type Conf struct {
	NamesrvAddr   string `yaml:"namesrv_addr"`
	BrokerAddr    string `yaml:"broker_addr"`
	RetryTimes    int    `yaml:"retry_times"`
	ProducerGroup string `yaml:"producer_group"`
}

func FromEnv() Conf {
	return Conf{
		NamesrvAddr:   env.String("ROCKETMQ_NAMESRV_ADDR"),
		BrokerAddr:    env.String("ROCKETMQ_BROKER_ADDR"),
		RetryTimes:    env.Int("ROCKETMQ_RETRY_TIMES", 1),
		ProducerGroup: env.String("ROCKETMQ_PRODUCER_GROUP", ""),
	}
}
