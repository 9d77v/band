package freecache

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/9d77v/band/pkg/stores/cache"
	"github.com/coocood/freecache"
)

func TestFreeCache_Get(t *testing.T) {
	c, _ := NewFreeCache(cache.FromEnv())
	key := "ddd"
	value := "www"
	err := c.Set(context.Background(), key, value, 10*time.Second)
	if err != nil {
		log.Panicln(err)
	}
	type fields struct {
		client *freecache.Cache
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
			r := &FreeCache{
				client: tt.fields.client,
				conf:   tt.fields.conf,
			}
			got, err := r.Get(tt.args.ctx, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("FreeCache.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FreeCache.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
