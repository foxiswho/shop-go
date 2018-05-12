package models

import (
	"time"
)

type BlogSyncQueue struct {
	QueueId    int       `json:"queue_id" xorm:"not null pk autoincr INT(11)"`
	BlogId     int       `json:"blog_id" xorm:"not null default 0 comment('本站博客id') INT(11)"`
	TypeId     int       `json:"type_id" xorm:"not null default 0 comment('类型') INT(11)"`
	Status     int       `json:"status" xorm:"not null default 0 comment('状态：0:待运行 10:失败 99:成功') TINYINT(3)"`
	TimeUpdate time.Time `json:"time_update" xorm:"default 'CURRENT_TIMESTAMP' comment('最后一次更新时间') TIMESTAMP"`
	TimeAdd    time.Time `json:"time_add" xorm:"default 'CURRENT_TIMESTAMP' comment('插入时间') TIMESTAMP"`
	Msg        string    `json:"msg" xorm:"comment('内容') VARCHAR(255)"`
	MapId      int       `json:"map_id" xorm:"not null default 0 comment('同步ID') INT(11)"`
}

//初始化
func NewBlogSyncQueue() *BlogSyncQueue {
	return new(BlogSyncQueue)
}
