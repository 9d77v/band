package freecache

import (
	"bytes"
	"context"
	"time"

	"encoding/gob"

	"github.com/9d77v/band/pkg/stores/cache"
	"github.com/coocood/freecache"
)

type FreeCache struct {
	client *freecache.Cache
	conf   cache.Conf
}

func NewFreeCache(conf cache.Conf) (*FreeCache, error) {
	client := freecache.NewCache(100 * 1024 * 1024)
	return &FreeCache{
		client: client,
		conf:   conf,
	}, nil
}

func (r *FreeCache) Get(ctx context.Context, key string) (string, error) {
	bs, err := r.client.Get([]byte(key))
	if err != nil {
		return "", err
	}
	var value string
	decoder := gob.NewDecoder(bytes.NewReader(bs))
	err = decoder.Decode(&value)
	return value, err
}
func (r *FreeCache) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(value)
	if err != nil {
		return err
	}
	return r.client.Set([]byte(key), buffer.Bytes(), int(expiration.Seconds()))
}
func (r *FreeCache) Del(ctx context.Context, key string) error {
	r.client.Del([]byte(key))
	return nil
}
func (r *FreeCache) GetDel(ctx context.Context, key string) (string, error) {
	value, err := r.Get(ctx, key)
	if err != nil {
		return "", err
	}
	return value, r.Del(ctx, key)
}
func (r *FreeCache) PTTL(ctx context.Context, key string) (time.Duration, error) {
	_, duration, err := r.client.GetWithExpiration([]byte(key))
	return time.Duration(duration) * time.Second, err
}
