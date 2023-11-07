package oss

import (
	"context"
	"io"
	"sync"
	"time"

	"github.com/9d77v/band/pkg/conf/oss/impl/alioss"
	"github.com/9d77v/band/pkg/conf/oss/impl/minio"
	"github.com/9d77v/band/pkg/conf/oss/ossconf"
)

type Oss interface {
	PresignedPutURL(ctx context.Context, objectName string, expires time.Duration, etag, mimeType string, size int64) (string, error)
	PresignedGetURL(ctx context.Context, objectName string, expires time.Duration) (string, error)
	StatObject(ctx context.Context, objectName string) (ossconf.ObjectInfo, error)
	PutObject(ctx context.Context, objectName string, reader io.Reader, objectSize int64, contentType string) error
	GetBucketName() string
	GetObjectName(prefix, fileName string) string
	GetObject(ctx context.Context, objectName string, opts ossconf.GetObjectOption) (io.Reader, error)
	GetStoragePath(prefix, fileName string) string
}

var (
	client Oss
	once   sync.Once
)

const (
	TypeMinio  = "minio"
	TypeAlioss = "alioss"
)

func NewOss(conf ossconf.Conf) Oss {
	once.Do(func() {
		switch conf.Type {
		case TypeAlioss:
			client = alioss.NewOSS(conf)
		default:
			client = minio.NewMinio(conf)
		}
	})
	return client
}
