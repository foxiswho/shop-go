package timer

import (
	"time"
)

//定时更新任务
func NewTimer(f func(), t time.Duration) {
	timer2 := time.NewTimer(t)
	go func() {
		//等触发时的信号
		<-timer2.C
		f()
	}()

}
