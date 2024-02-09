package cache_factory

import (
	"sync"

	"github.com/9d77v/band/pkg/stores/cache"
	"github.com/9d77v/band/pkg/stores/cache/impl/freecache"
	"github.com/9d77v/band/pkg/stores/cache/impl/redis"
)

var (
	client cache.Cache
	once   sync.Once
)

const (
	CacheTypeLocal = "local"
	CacheTypeRedis = "redis"
)

func NewCache(conf cache.Conf) (cache.Cache, error) {
	var client cache.Cache
	var err error
	switch conf.Type {
	case CacheTypeRedis:
		client, err = redis.NewRedis(conf)
	default:
		client, err = freecache.NewFreeCache(conf)
	}
	return client, err
}

func CacheSingleton(conf cache.Conf) (cache.Cache, error) {
	var err error
	once.Do(func() {
		client, err = NewCache(conf)
	})
	return client, err
}
