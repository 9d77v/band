package consumer

import "github.com/9d77v/band/pkg/env"

type Conf struct {
	Type         string   `yaml:"type"`
	Addrs        []string `yaml:"addrs"`
	GroupName    string   `yaml:"group_name"`
	ConsumeModel int      `yaml:"consume_model"`
}

func FromEnv() Conf {
	return Conf{
		Type:         env.String("CONSUMER_TYPE"),
		Addrs:        env.StringArray("CONSUMER_ADDRS", ","),
		GroupName:    env.String("CONSUMER_GROUP_NAME", ""),
		ConsumeModel: env.Int("CONSUMER_CONSUME_MODEL", 1),
	}
}
