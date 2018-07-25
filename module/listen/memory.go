package listen

import (
	"github.com/foxiswho/shop-go/module/cache"
	cache2 "github.com/foxiswho/shop-go/middleware/cache"
	"time"
	"github.com/foxiswho/shop-go/consts/cache/cacheCache"
	"github.com/foxiswho/shop-go/module/cache/cacheMemory"
	"github.com/foxiswho/shop-go/module/log"
	"strconv"
)

var (
	Listen_Memory_Second = time.Hour * 24 * 365
)
//监听缓存是否biang，如果变化则更新缓存
func ListenMemory() {
	client := cache.ClientRedis()
	redis := client.(*cache2.RedisStore)
	//读取缓存，获取要更新的时间戳
	systemCacheTime, err := redis.HGetAll(cacheCache.System_Cache_Memory_Sync_Time_Stamp)
	if err != nil {
		log.Debugf("ListenMemory HGetAll error: %v", err)
		panic(err)
	}
	//时间戳是否 有值
	if systemCacheTime != nil && len(systemCacheTime) > 0 {
		memoryCacheTime := make(map[string]int)
		//获取内存中 该键最后更新时间
		err := cacheMemory.MemoryGet(cacheCache.System_Cache_Memory_Sync, &memoryCacheTime)
		if err != nil {
			log.Debugf("Listen Memory in cacheMemory error: %v", err)
			memoryCacheTime = make(map[string]int)
		}
		updateFields := []string{}
		for key, val := range systemCacheTime {
			intVal, _ := strconv.Atoi(val)
			if intVal > 0 {
				if _, ok := memoryCacheTime[key]; ok {
					if memoryCacheTime[key] < intVal {
						//赋值时间戳
						memoryCacheTime[key] = intVal
						//加入要更新的 标志
						updateFields = append(updateFields, key)
					}
				} else {
					//赋值时间戳
					memoryCacheTime[key] = intVal
					//加入要更新的 标志
					updateFields = append(updateFields, key)
				}
			}
		}
		// 更新指定缓存到内存中
		if len(systemCacheTime) > 0 {
			log.Debugf("updated cache fields : %v ", updateFields)
			cacheMemory.MemoryUpdateByCache(updateFields, memoryCacheTime)
		}
	}
}
