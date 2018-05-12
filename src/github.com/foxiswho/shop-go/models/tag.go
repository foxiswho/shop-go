package models

import (
	"time"
)

type Tag struct {
	TagId   int       `json:"tag_id" xorm:"not null pk autoincr INT(11)"`
	Name    string    `json:"name" xorm:"comment('名称') CHAR(50)"`
	TimeAdd time.Time `json:"time_add" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
}

//初始化
func NewTag() *Tag {
	return new(Tag)
}
