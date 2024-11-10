package cache

import (
	"context"
	"time"
)

type Cache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Del(ctx context.Context, key string) error
	GetDel(ctx context.Context, key string) (string, error)
	PTTL(ctx context.Context, key string) (time.Duration, error)
}
