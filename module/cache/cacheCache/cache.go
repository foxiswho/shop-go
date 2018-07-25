package cacheCache

import (
	"github.com/foxiswho/shop-go/models/setting/siteSetting"
	"github.com/foxiswho/shop-go/module/cache"
	"time"
	"github.com/foxiswho/shop-go/consts/cache/cacheCache"
	"github.com/foxiswho/shop-go/consts/cache/cacheMemory"
	"github.com/foxiswho/shop-go/module/log"
)

var (
	//缓存时间
	Cache_Second = time.Hour * 24 * 365
)

//获取 缓存
func Get(key string, value interface{}) (error) {
	return cache.ClientRedis().Get(key, value)
}

//设置 缓存
func Set(key string, value interface{}, expire time.Duration) (error) {
	return cache.ClientRedis().Set(key, value, expire)
}

//删除 缓存
func Del(key string) (error) {
	return cache.ClientRedis().Delete(key)
}

func LoadOneCache() {
	redis := cache.ClientRedisStore()
	if redis.HExists(cacheCache.System_Cache, cacheMemory.SiteSetting) {
		//
		log.Debugf("cacheCache.System_Cache cacheMemory.SiteSetting Find")
	} else {
		log.Debugf("cacheCache.System_Cache cacheMemory.SiteSetting SET")
		site := &siteSetting.Site{}
		site.SiteName = "SHOP"
		err := redis.HSet(cacheCache.System_Cache, cacheMemory.SiteSetting, site, 0)
		log.Debugf("cacheCache.System_Cache cacheMemory.SiteSetting SET RESULT", err)
	}
}
