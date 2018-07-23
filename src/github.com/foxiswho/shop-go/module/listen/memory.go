package listen

import (
	"github.com/foxiswho/shop-go/module/cache"
	cache2 "github.com/foxiswho/shop-go/middleware/cache"
	"time"
	"github.com/foxiswho/shop-go/consts/cache/cacheCache"
	"github.com/foxiswho/shop-go/module/cache/cacheMemory"
	"github.com/foxiswho/shop-go/module/log"
)

var (
	Listen_Memory_Second = time.Hour * 24 * 365
)
//监听更新缓存
func ListenMemory() {
	client := cache.ClientRedis()
	redis := client.(*cache2.RedisStore)
	//读取缓存中，同步的时间戳
	arr, err := redis.HGetAll(cacheCache.System_Cache_Memory_Sync)
	if err != nil {
		return nil, err
	}
	if arr != nil && len(arr) > 0 {
		memory_cache := make(map[string]int)
		//获取内存中 该键最后更新时间
		err := cacheMemory.MemoryGet(cacheCache.System_Cache_Memory_Sync, &memory_cache)
		if err != nil {
			log.Debugf("Listen Memory in cacheMemory error: %v", err)
			memory_cache = make(map[string]int)
		}
		update_fields := []string{}
		for key, val := range arr {
			if val != nil {
				i := val.(int)
				if _, ok := memory_cache["xx"]; ok {
					if memory_cache["xx"] < i {
						//更新指定缓存
						update_fields = append(update_fields, key)
					}
				} else {
					//更新指定缓存
					update_fields = append(update_fields, key)
				}
			}
		}
		cacheMemory.MemoryUpdateByCache(update_fields, memory_cache)
	}

}
