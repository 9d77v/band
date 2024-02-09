package redis

import (
	"context"
	"log"
	"time"

	"github.com/9d77v/band/pkg/stores/cache"
	redis "github.com/redis/go-redis/v9"
)

type Redis struct {
	client redis.UniversalClient
	conf   cache.Conf
}

func NewRedis(conf cache.Conf) (*Redis, error) {
	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    conf.Addrs,
		Password: conf.Password,
		DB:       conf.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	log.Println("connected to redis:", client)
	return &Redis{
		client: client,
		conf:   conf,
	}, nil
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *Redis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}
func (r *Redis) Del(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}
func (r *Redis) GetDel(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}
func (r *Redis) PTTL(ctx context.Context, key string) (time.Duration, error) {
	return r.client.PTTL(ctx, key).Result()
}
