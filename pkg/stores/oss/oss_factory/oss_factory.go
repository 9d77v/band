package oss_factory

import (
	"log"
	"sync"

	"github.com/9d77v/band/pkg/stores/oss"
	"github.com/9d77v/band/pkg/stores/oss/impl/alioss"
	"github.com/9d77v/band/pkg/stores/oss/impl/minio"
)

var (
	client oss.Oss
	once   sync.Once
)

const (
	TypeMinio  = "minio"
	TypeAlioss = "alioss"
)

func NewOss(conf oss.Conf) oss.Oss {
	var client oss.Oss
	switch conf.Type {
	case TypeAlioss:
		client = alioss.NewOSS(conf)
	default:
		client = minio.NewMinio(conf)
	}
	log.Println("connected to oss:", client)
	return client
}

func OssSingleton(conf oss.Conf) oss.Oss {
	once.Do(func() {
		client = NewOss(conf)
	})
	return client
}
