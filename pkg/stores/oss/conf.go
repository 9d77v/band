package oss

import "github.com/9d77v/band/pkg/env"

type Conf struct {
	Type           string `yaml:"type"`
	Addr           string `yaml:"addr"`
	ExternalAddr   string `yaml:"external_addr"`
	Secure         bool   `yaml:"secure"`
	AccessKey      string `yaml:"access_key"`
	SecretKey      string `yaml:"secret_key"`
	ExpiryDuration int    `yaml:"expiry_duration"`
	BucketName     string `yaml:"bucket_name"`
}

func FromEnv() Conf {
	return Conf{
		Type:         env.String("OSS_TYPE"),
		Addr:         env.String("OSS_ADDR", "localhost:9000"),
		ExternalAddr: env.String("OSS_EXTERNAL_ADDR", "http://localhost:9000"),
		Secure:       env.Bool("OSS_SECURE"),
		AccessKey:    env.String("OSS_ACCESS_KEY", "minio"),
		SecretKey:    env.String("OSS_SECRET_KEY", "minio123"),
		BucketName:   env.String("OSS_BUCKET_NAME"),
	}
}
