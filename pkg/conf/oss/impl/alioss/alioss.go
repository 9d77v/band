package alioss

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/9d77v/band/pkg/conf/oss/ossconf"
	ossSDK "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type OSS struct {
	client *ossSDK.Bucket
	conf   ossconf.Conf
}

func NewOSS(conf ossconf.Conf) *OSS {
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

func (m *OSS) PresignedPutURL(ctx context.Context, objectName string, expires time.Duration, etag, mimeType string, size int64) (string, error) {
	options := []ossSDK.Option{
		ossSDK.ContentType(mimeType),
	}
	return m.client.SignURL(objectName, ossSDK.HTTPPut, int64(expires.Seconds()), options...)
}

func (m *OSS) PresignedGetURL(ctx context.Context, objectName string, expires time.Duration) (string, error) {
	return m.client.SignURL(objectName, ossSDK.HTTPGet, int64(expires.Seconds()))
}

func (m *OSS) StatObject(ctx context.Context, objectName string) (ossconf.ObjectInfo, error) {
	obj, err := m.client.GetObjectMeta(objectName)
	size, _ := strconv.ParseInt(obj.Get("Content-Length"), 10, 64)
	return ossconf.ObjectInfo{
		Size: size,
	}, err
}

func (m *OSS) GetBucketName() string {
	return m.conf.BucketName
}

func (m *OSS) GetObjectName(prefix, fileName string) string {
	return fmt.Sprintf("%s/%s", prefix, fileName)
}

func (m *OSS) GetStoragePath(prefix, fileName string) string {
	return fmt.Sprintf("/%s/%s", m.conf.BucketName, m.GetObjectName(prefix, fileName))
}
