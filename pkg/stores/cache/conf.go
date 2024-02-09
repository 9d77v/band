package cache

import "github.com/9d77v/band/pkg/env"

type Conf struct {
	Type     string   `yaml:"type"`
	Addrs    []string `yaml:"addrs"`
	Password string   `yaml:"password"`
	DB       int      `yaml:"db"`
}

func FromEnv() Conf {
	return Conf{
		Type:     env.String("CACHE_TYPE", "LOCAL"),
		Addrs:    env.StringArray("CACHE_ADDRESS", ",", "localhost:6379"),
		Password: env.String("CACHE_PASSWORD"),
		DB:       env.Int("CACHE_DB"),
	}
}
