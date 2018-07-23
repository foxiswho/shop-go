package crontab

import (
	"github.com/robfig/cron"
	"github.com/foxiswho/shop-go/module/log"
	"time"
	"github.com/foxiswho/shop-go/module/cache/cacheMemory"
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
	log.Debugf("cron closed:", time.Now())
}

func Task() {
	if c == nil {
		newCron()
	}
	spec := "*/5 * * * * ?"
	c.AddFunc(spec, func() {
		//只有内存数据加载过了，才执行此任务
		if cacheMemory.Is_Load_Once {
			log.Debugf("cron time:", time.Now())
		}
	})
	c.Start()
	log.Debugf("cron start:", time.Now())
}
