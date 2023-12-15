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
}

func FromEnv(serviceName string) Conf {
	return Conf{
		AppName:     env.String("APP_NAME"),
		ServiceName: serviceName,
		ServerHost:  env.String("SERVER_HOST"),
		ServerPort:  uint64(env.Int("SERVER_PORT")),
		EtcdAddress: env.String("ETCD_ADDRESS"),
	}
}

func RPCFromEnv(serviceName string) Conf {
	return Conf{
		AppName:     env.String("APP_NAME"),
		ServiceName: serviceName,
		ServerHost:  network.GetNetworkIp(),
		ServerPort:  network.GetRandomPort(),
		EtcdAddress: env.String("ETCD_ADDRESS", "http://localhost:2379"),
	}
}
