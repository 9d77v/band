package minio

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
	cr "github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/9d77v/band/pkg/conf/oss/ossconf"
)

type Minio struct {
	client *minio.Client
	conf   ossconf.Conf
}

func NewMinio(conf ossconf.Conf) *Minio {
	client, err := minio.New(conf.Addr, &minio.Options{
		Creds:  cr.NewStaticV4(conf.AccessKey, conf.SecretKey, ""),
		Secure: conf.Secure,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return &Minio{
		client: client,
		conf:   conf,
	}
}

func (m *Minio) GetSTSAssumeRole(uid string, policyPattern string) (cr.Value, error) {
	var stsOpts cr.STSAssumeRoleOptions
	stsOpts.AccessKey = m.conf.AccessKey
	stsOpts.SecretKey = m.conf.SecretKey
	stsOpts.Policy = fmt.Sprintf(policyPattern, uid)
	if m.conf.ExpiryDuration != 0 {
		stsOpts.DurationSeconds = m.conf.ExpiryDuration
	}
	var err error
	cred, err := cr.NewSTSAssumeRole(m.conf.Addr, stsOpts)
	if err != nil {
		return cr.Value{}, err
	}
	return cred.Get()
}

func (m *Minio) PresignedPutURL(ctx context.Context, objectName string, expires time.Duration, etag, mimeType string, size int64) (string, error) {
	u, err := m.client.PresignHeader(ctx, http.MethodPut, m.conf.BucketName, objectName, expires, url.Values{}, http.Header{"x-amz-content-sha256": []string{etag}})
	if err != nil {
		return "", err
	}
	return u.String(), err
}

func (m *Minio) PresignedGetURL(ctx context.Context, objectName string, expires time.Duration) (string, error) {
	u, err := m.client.PresignedGetObject(ctx, m.conf.BucketName, objectName, expires, url.Values{})
	if err != nil {
		return "", err
	}
	return u.String(), err
}

func (m *Minio) PutObject(ctx context.Context, objectName string, reader io.Reader, objectSize int64, contentType string) error {
	_, err := m.client.PutObject(ctx, m.conf.BucketName, objectName, reader, objectSize, minio.PutObjectOptions{ContentType: contentType})
	return err
}

func (m *Minio) GetObject(ctx context.Context, objectName string, opts ossconf.GetObjectOption) (io.Reader, error) {
	options := minio.GetObjectOptions{}
	bytesRange := opts.GetRange()
	if bytesRange != nil {
		err := options.SetRange(bytesRange.Start, bytesRange.End)
		log.Println("set range failed", err)
	}
	return m.client.GetObject(ctx, m.conf.BucketName, objectName, options)
}

func (m *Minio) StatObject(ctx context.Context, objectName string) (ossconf.ObjectInfo, error) {
	obj, err := m.client.StatObject(ctx, m.conf.BucketName, objectName, minio.GetObjectOptions{})
	return ossconf.ObjectInfo{
		Size: obj.Size,
	}, err
}

func (m *Minio) GetBucketName() string {
	return m.conf.BucketName
}

func (m *Minio) GetObjectName(prefix, fileName string) string {
	return fmt.Sprintf("%s/%s", prefix, fileName)
}

func (m *Minio) GetStoragePath(prefix, fileName string) string {
	return fmt.Sprintf("/%s/%s", m.conf.BucketName, m.GetObjectName(prefix, fileName))
}
