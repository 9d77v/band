package grpc

import "github.com/9d77v/band/pkg/env"

type Conf struct {
	AppName     string `yaml:"app_name"`
	ServiceName string
	EtcdAddress string `yaml:"etcd_address"`
}

func FromEnv(serviceName string) Conf {
	return Conf{
		AppName:     env.String("APP_NAME"),
		EtcdAddress: env.String("ETCD_ADDRESS"),
		ServiceName: serviceName,
	}
}
