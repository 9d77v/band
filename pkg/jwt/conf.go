package jwt

import "github.com/9d77v/band/pkg/env"

type Conf struct {
	SecretKey  string `yaml:"secret_key"`
	ExpireTime int64  `yaml:"expire_time"`
	Issuer     string `yaml:"issuer"`
}

func FromEnv() Conf {
	return Conf{
		SecretKey:  env.String("JWT_SECRET_KEY"),
		ExpireTime: int64(env.Int("JWT_EXPIRE_TIME")),
		Issuer:     env.String("JWT_ISSUER"),
	}
}
