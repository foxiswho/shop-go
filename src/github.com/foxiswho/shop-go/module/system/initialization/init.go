package initialization

import (
	"github.com/foxiswho/shop-go/module/cache/cacheCache"
	"github.com/foxiswho/shop-go/module/cache/cacheMemory"
	"sync"
	"github.com/foxiswho/shop-go/module/log"
)

var (
	//只执行一次
	once sync.Once
)
//初始化
func InitSystem() {
	//缓存
	cacheCache.LoadOneCache()
	//只执行一次
	once.Do(onces)
}

//只执行一次
func onces() {
	log.Debugf("sync.Once 只加载一次缓存 cacheMemory.LoadOneCache,")
	//内存缓存
	cacheMemory.LoadOneCache()
}
