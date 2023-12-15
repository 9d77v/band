package redis

import "github.com/9d77v/band/pkg/env"

type Conf struct {
	Addrs    []string `yaml:"addrs"`
	Password string   `yaml:"password"`
	DB       int      `yaml:"db"`
}

func FromEnv() Conf {
	return Conf{
		Addrs:    env.StringArray("REDIS_ADDRESS", ",", "localhost:6379"),
		Password: env.String("REDIS_PASSWORD"),
		DB:       env.Int("REDIS_DB"),
	}
}
