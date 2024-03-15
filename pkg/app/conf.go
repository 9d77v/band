package app

import (
	"github.com/9d77v/band/pkg/env"
	"github.com/9d77v/band/pkg/network"
)

type Conf struct {
	AppName     string `yaml:"app_name"`
	ServiceName string
	ServerHost  string `yaml:"server_host"`
	ServerPort  uint64 `yaml:"server_port"`
	EtcdAddress string `yaml:"etcd_address"`
	MaxConns    int    `yaml:"max_conns"`
}

func FromEnv(serviceName string) Conf {
	return Conf{
		AppName:     env.String("APP_NAME"),
		ServiceName: serviceName,
		ServerHost:  env.String("SERVER_HOST"),
		ServerPort:  uint64(env.Int("SERVER_PORT")),
		EtcdAddress: env.String("ETCD_ADDRESS"),
		MaxConns:    env.Int("MAX_CONNS", 10000),
	}
}

func RPCFromEnv(serviceName string) Conf {
	return Conf{
		AppName:     env.String("APP_NAME"),
		ServiceName: serviceName,
		ServerHost:  network.GetNetworkIp(),
		ServerPort:  network.GetRandomPort(),
		EtcdAddress: env.String("ETCD_ADDRESS", "http://localhost:2379"),
		MaxConns:    env.Int("MAX_CONNS", 10000),
	}
}
