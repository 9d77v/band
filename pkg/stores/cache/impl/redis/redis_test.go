package redis

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/9d77v/band/pkg/stores/cache"
	redis "github.com/redis/go-redis/v9"
)

func TestRedis_Get(t *testing.T) {
	conf := cache.FromEnv()
	conf.Type = "redis"
	c, _ := NewRedis(conf)
	key := "ddd"
	value := "www"
	err := c.Set(context.Background(), key, value, 10*time.Second)
	if err != nil {
		log.Panicln(err)
	}
	type fields struct {
		client redis.UniversalClient
		conf   cache.Conf
	}
	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{"test cache get", fields{c.client, c.conf}, args{context.Background(),
			key}, value, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Redis{
				client: tt.fields.client,
				conf:   tt.fields.conf,
			}
			got, err := r.Get(tt.args.ctx, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Redis.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Redis.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
