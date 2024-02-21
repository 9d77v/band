package oss

import (
	"context"
	"io"
	"time"
)

type Oss interface {
	PresignedPutURL(ctx context.Context, objectName string, expires time.Duration, etag, mimeType string, size int64) (string, error)
	PresignedGetURL(ctx context.Context, objectName string, expires time.Duration) (string, error)
	StatObject(ctx context.Context, objectName string) (ObjectInfo, error)
	PutObject(ctx context.Context, objectName string, reader io.Reader, objectSize int64, contentType string) error
	GetBucketName() string
	GetObject(ctx context.Context, objectName string, opts GetObjectOption) (io.Reader, error)
	DeleteObject(ctx context.Context, objectName string) error
}
