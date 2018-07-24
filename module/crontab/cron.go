package crontab

import (
	"github.com/robfig/cron"
	"github.com/foxiswho/shop-go/module/log"
	"github.com/foxiswho/shop-go/module/cache/cacheMemory"
	"github.com/foxiswho/shop-go/util/datetime"
	"github.com/foxiswho/shop-go/module/listen"
)

var c *cron.Cron

func newCron() {
	c = cron.New()
}
func TaskClosed() {
	if c != nil {
		c.Stop()
	}
	c = nil
	log.Debugf("cron closed:", datetime.Now())
}

func Task() {
	if c == nil {
		newCron()
	}
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {
		//只有内存数据加载过了，才执行此任务
		if cacheMemory.Is_Load_Once {
			//开始执行定时任务
			//定时更新内存缓存
			listen.ListenMemory()
		} else {
			log.Debugf("cron waite:", datetime.Now())
		}
	})
	c.Start()
	//log.Debugf("cron start:", time.Now())
}
