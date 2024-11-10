package alioss

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	"github.com/9d77v/band/pkg/stores/oss"
	ossSDK "github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
)

type OSS struct {
	client *ossSDK.Client
	conf   oss.Conf
}

func NewOSS(conf oss.Conf) *OSS {
	provider := credentials.NewStaticCredentialsProvider(conf.AccessKey, conf.SecretKey)
	cfg := ossSDK.LoadDefaultConfig().
		WithCredentialsProvider(provider).
		WithRegion(conf.Region)
	return &OSS{
		client: ossSDK.NewClient(cfg),
		conf:   conf,
	}
}

func (m *OSS) ExternalAddr() string {
	return m.conf.ExternalAddr
}

func (m *OSS) PresignedPutURL(ctx context.Context, objectName string, expires time.Duration, etag, mimeType string, size int64) (string, error) {
	result, err := m.client.Presign(
		ctx,
		&ossSDK.PutObjectRequest{
			Bucket:        &m.conf.BucketName,
			Key:           &objectName,
			ContentLength: &size,
			ContentType:   &mimeType,
		},
		ossSDK.PresignExpires(expires),
	)
	if err != nil {
		return "", err
	}
	return result.URL, nil
}

func (m *OSS) PresignedGetURL(ctx context.Context, objectName string, expires time.Duration) (string, error) {
	result, err := m.client.Presign(
		ctx,
		&ossSDK.GetObjectRequest{
			Bucket: &m.conf.BucketName,
			Key:    &objectName,
		},
		ossSDK.PresignExpires(expires),
	)
	if err != nil {
		return "", err
	}
	return result.URL, nil
}

func (m *OSS) PutObject(ctx context.Context, objectName string, reader io.Reader, objectSize int64, contentType string) error {
	_, err := m.client.PutObject(ctx, &ossSDK.PutObjectRequest{
		Bucket:        &m.conf.BucketName,
		Key:           &objectName,
		Body:          reader,
		ContentType:   &contentType,
		ContentLength: &objectSize,
	})
	return err
}

func (m *OSS) GetObject(ctx context.Context, objectName string, opts oss.GetObjectOption) (io.Reader, error) {
	bytesRange := opts.GetRange()
	var rangeStr *string
	if bytesRange != nil {
		obj, err := m.StatObject(ctx, objectName)
		if err != nil {
			return nil, err
		}
		str := fmt.Sprintf("%d~%d/%d", bytesRange.Start, bytesRange.End, obj.Size)
		rangeStr = &str
	}
	res, err := m.client.GetObject(ctx, &ossSDK.GetObjectRequest{
		Bucket: &m.conf.BucketName,
		Key:    &objectName,
		Range:  rangeStr,
	})
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}

func (m *OSS) StatObject(ctx context.Context, objectName string) (oss.ObjectInfo, error) {
	obj, err := m.client.GetObjectMeta(ctx, &ossSDK.GetObjectMetaRequest{
		Bucket: &m.conf.BucketName,
		Key:    &objectName,
	})
	if err != nil {
		return oss.ObjectInfo{}, err
	}
	return oss.ObjectInfo{
		Size: obj.ContentLength,
	}, err
}

func (m *OSS) GetBucketName() string {
	return m.conf.BucketName
}

func (m *OSS) DeleteObject(ctx context.Context, objectName string) error {
	_, err := m.client.DeleteObject(ctx, &ossSDK.DeleteObjectRequest{
		Bucket: &m.conf.BucketName,
		Key:    &objectName,
	})
	return err
}
