package cache

import (
	ec "github.com/foxiswho/shop-go/middleware/cache"
	. "github.com/foxiswho/shop-go/module/conf"
	"time"
)

var client map[string]ec.CacheStore

//新客户端
func NewClient(cacheStore string) ec.CacheStore {
	var store ec.CacheStore

	switch cacheStore {
	case MEMCACHED:
		store = ec.NewMemcachedStore([]string{Conf.Memcached.Server}, time.Hour)
	case REDIS:
		store = ec.NewRedisCache(Conf.Redis.Server, Conf.Redis.Pwd, DefaultExpiration)
	default:
		store = ec.NewInMemoryStore(time.Hour)
	}
	return store
}
func clientInit(cacheStore string) ec.CacheStore {
	if client == nil {
		client = make(map[string]ec.CacheStore)
	}
	if _, ok := client[cacheStore]; ok {
		return client[cacheStore]
	}
	client[cacheStore] = NewClient(cacheStore)
	return client[cacheStore]
}

// 缓存
func Client() ec.CacheStore {
	return clientInit(Conf.CacheStore)
}

// 内存 缓存
func ClientMemory() ec.CacheStore {
	return clientInit(IN_MEMORY)
}

// Redis 缓存
func ClientRedis() ec.CacheStore {
	return clientInit(REDIS)
}

// Redis 缓存
func ClientRedisStore() *ec.RedisStore {
	return clientInit(REDIS).(*ec.RedisStore)
}
