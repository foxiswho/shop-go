package orm

import (
	"github.com/hb-go/echo-web/module/log"
)

type Logger struct {
}

// Print format & print log
func (logger Logger) Print(values ...interface{}) {
	// @TODO
	// 日志格式化解析
	log.Debugf("orm log:%v \n", values)
}
