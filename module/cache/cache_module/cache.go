package cache_module

import (
	"github.com/foxiswho/shop-go/models/setting/site_setting"
	"github.com/foxiswho/shop-go/module/cache"
	"time"
	"github.com/foxiswho/shop-go/consts/cache/cache_consts"
	"github.com/foxiswho/shop-go/consts/cache/memory_consts"
	"github.com/foxiswho/shop-go/module/log"
	"github.com/foxiswho/shop-go/module/timer"
)

var (
	//缓存时间
	Cache_Second               = time.Hour * 24 * 365
	System_Cache_Sync_Second_2 = time.Minute * 5
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
	if redis.HExists(cache_consts.System_Cache, memory_consts.SiteSetting) {
		//
		log.Debugf("cacheCache.System_Cache cacheMemory.SiteSetting Find")
	} else {
		log.Debugf("cacheCache.System_Cache cacheMemory.SiteSetting SET")
		site := &site_setting.Site{}
		site.SiteName = "SHOP"
		err := redis.HSet(cache_consts.System_Cache, memory_consts.SiteSetting, site, 0)
		log.Debugf("cacheCache.System_Cache cacheMemory.SiteSetting SET RESULT", err)
	}
}

//获取 同步缓存
func SystemGet(key, field string, value interface{}) (error) {
	return cache.ClientRedisStore().HGet(key, field, value)
}

//设置 同步缓存
func SystemSet(key, field string, value interface{}, expire time.Duration) (error) {
	err := cache.ClientRedisStore().HSet(key, field, value, expire)
	//定时器更新二级缓存
	timer.NewTimer(func(key, field string) {
		updateSystemSet2(key, field)
	}, System_Cache_Sync_Second_2)
	return err
}

//获取 同步二级缓存
func SystemGet2(key, field string, value interface{}) (error) {
	return cache.ClientRedisStore().HGet(key+cache_consts.System_Cache_Sync_Level2_Mark, field, value)
}

//设置 同步二级缓存
func SystemSet2(key, field string, value interface{}, expire time.Duration) (error) {
	return cache.ClientRedisStore().HSet(key+cache_consts.System_Cache_Sync_Level2_Mark, field, value, expire)
}

//更新二级缓存
func updateSystemSet2(key, field string) {
	var value interface{}
	//获取1级缓存
	err := SystemGet(key, field, value)
	if err != nil {
		log.Debugf("updateSystemSet2:%v %v error:%v", key, field, err)
	} else {
		//更新二级缓存
		err := SystemSet2(key, field, value, Cache_Second)
		if err != nil {
			log.Debugf("updateSystemSet2:%v %v error:%v", key, field, err)
		} else {
			log.Debugf("updateSystemSet2 success:%v %v value:%v", key, field, value)
		}
	}
}
