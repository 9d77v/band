package alioss

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	"github.com/9d77v/band/pkg/stores/oss"
	ossSDK "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type OSS struct {
	client *ossSDK.Bucket
	conf   oss.Conf
}

func NewOSS(conf oss.Conf) *OSS {
	client, err := ossSDK.New(fmt.Sprintf("https://%s", conf.Addr),
		conf.AccessKey, conf.SecretKey)
	if err != nil {
		log.Fatalln(err)
	}
	bucket, err := client.Bucket(conf.BucketName)
	if err != nil {
		log.Fatalln(err)
	}
	return &OSS{
		client: bucket,
		conf:   conf,
	}
}

func (m *OSS) ExternalAddr() string {
	return m.conf.ExternalAddr
}

func (m *OSS) PresignedPutURL(ctx context.Context, objectName string, expires time.Duration, etag, mimeType string, size int64) (string, error) {
	options := []ossSDK.Option{
		ossSDK.ContentType(mimeType),
	}
	return m.client.SignURL(objectName, ossSDK.HTTPPut, int64(expires.Seconds()), options...)
}

func (m *OSS) PresignedGetURL(ctx context.Context, objectName string, expires time.Duration) (string, error) {
	return m.client.SignURL(objectName, ossSDK.HTTPGet, int64(expires.Seconds()))
}

func (m *OSS) PutObject(ctx context.Context, objectName string, reader io.Reader, objectSize int64, contentType string) error {
	return m.client.PutObject(objectName, reader)
}

func (m *OSS) GetObject(ctx context.Context, objectName string, opts oss.GetObjectOption) (io.Reader, error) {
	ossOptions := []ossSDK.Option{}
	bytesRange := opts.GetRange()
	if bytesRange != nil {
		ossOptions = append(ossOptions,
			ossSDK.Range(bytesRange.Start, bytesRange.End))
	}
	body, err := m.client.GetObject(objectName, ossOptions...)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(body)
	body.Close()
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}

func (m *OSS) StatObject(ctx context.Context, objectName string) (oss.ObjectInfo, error) {
	obj, err := m.client.GetObjectMeta(objectName)
	size, _ := strconv.ParseInt(obj.Get("Content-Length"), 10, 64)
	return oss.ObjectInfo{
		Size: size,
	}, err
}

func (m *OSS) GetBucketName() string {
	return m.conf.BucketName
}

func (m *OSS) DeleteObject(ctx context.Context, objectName string) error {
	return m.client.DeleteObject(objectName)
}
