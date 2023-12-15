package producer

import "github.com/9d77v/band/pkg/env"

type Conf struct {
	Type       string   `yaml:"type"`
	Addrs      []string `yaml:"addrs"`
	RetryTimes int      `yaml:"retry_times"`
	GroupName  string   `yaml:"group_name"`
}

func FromEnv() Conf {
	return Conf{
		Type:       env.String("PRODUCER_TYPE"),
		Addrs:      env.StringArray("PRODUCER_ADDRS", ","),
		RetryTimes: env.Int("PRODUCER_RETRY_TIMES", 1),
		GroupName:  env.String("PRODUCER_GROUP_NAME", ""),
	}
}
