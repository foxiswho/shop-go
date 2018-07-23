package cacheMemory

import (
	"github.com/foxiswho/shop-go/module/cache"
	"time"
	"github.com/foxiswho/shop-go/consts/cache/cacheMemory"
	cache2 "github.com/foxiswho/shop-go/middleware/cache"
	"github.com/foxiswho/shop-go/consts/cache/cacheCache"
	"fmt"
	"github.com/foxiswho/shop-go/module/log"
	"github.com/foxiswho/shop-go/util/conv"
)

var (
	//只执行一次
	Is_Load_Once = false
	//缓存时间
	Memory_Second = time.Hour * 24 * 365
)

//初始化加载缓存
func LoadOneCache() {
	Is_Load_Once = true
	//只执行一次
	err := loadOneCache()
	if err != nil {
		log.Debugf("LoadOneCache error: %v", err)
	}
}

func loadOneCache() error {
	client := cache.ClientRedis()
	redis := client.(*cache2.RedisStore)
	//获取所有键值
	fields := MemoryFields()
	//获取所有系统缓存
	arr, err := redis.HGetAll(cacheCache.System_Cache)
	if err != nil {
		return err
	}
	fmt.Println("HGetAll System_Cache", arr)
	//获取系统缓存最后更新时间戳 读取缓存中，同步的时间戳
	arrSystem, err := redis.HGetAll(cacheCache.System_Cache_Memory_Sync)
	if err != nil {
		return err
	}
	fmt.Println("HGetAll", arrSystem)
	if arr != nil && len(arr) > 0 {
		memory := make(map[string]int)
		for _, key := range fields {
			if _, ok := arr[key]; ok {
				//设置 缓存
				MemorySet(key, arr[key], Memory_Second)
				//memory[key]=arrSystem["XX"]
				if _, is := arrSystem[key]; is {
					i, _ := conv.ObjToInt(arrSystem[key])
					memory[key] = i
				} else {
					memory[key] = 0
				}
			}
			var tmp interface{}
			err := MemoryGet(key, &tmp)
			log.Debugf("MemoryGet %v => %v |||err=> %v ", key, tmp, err)
		}
		// 存储 更新时间戳
		err = MemorySet(cacheCache.System_Cache_Memory_Sync, memory, Memory_Second)
		if err != nil {
			log.Debugf("Listen Memory in cacheMemory error: %v", err)
		}
	}
	return nil
}

//从cahce，更新指定缓存到内存中
func MemoryUpdateByCache(fields []string, memory_cache map[string]int) {
	client := cache.ClientRedis()
	redis := client.(*cache2.RedisStore)
	find := []string{}
	//读取缓存中，同步的时间戳
	arrSystem, err := redis.HGetAll(cacheCache.System_Cache_Memory_Sync)
	if err != nil {
		log.Debugf("HGetAll error: %v", err)
	}
	fmt.Println(arrSystem)
	for _, key := range fields {
		is_find := false
		//查找是否在数组中
		for _, k := range find {
			if k == key {
				is_find = true
				break;
			}
		}
		//如果不存在那么 进行更新
		if is_find == false {
			val := 0
			//获取缓存
			redis.HGet(cacheCache.System_Cache, key, val)
			//更新 缓存
			MemorySet(key, val, Memory_Second)
			// 存储 更新时间戳
			err = MemorySet(cacheCache.System_Cache_Memory_Sync, memory_cache, Memory_Second)
			if err != nil {
				log.Debugf("Listen Memory in cacheMemory error: %v", err)
			}
		}
	}
}

//所有键名
func MemoryFields() []string {
	return []string{cacheMemory.SiteSetting}
}

//获取 缓存
func MemoryGet(key string, value interface{}) (error) {
	return cache.ClientMemory().Get(key, value)
}

//设置 缓存
func MemorySet(key string, value interface{}, expire time.Duration) (error) {
	return cache.ClientMemory().Set(key, value, expire)
}

//删除 缓存
func MemoryDel(key string) (error) {
	return cache.ClientMemory().Delete(key)
}
