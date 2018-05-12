package models

import (
	"time"
)

type BlogTag struct {
	TagId   int       `json:"tag_id" xorm:"not null pk autoincr INT(11)"`
	Name    string    `json:"name" xorm:"comment('名称') CHAR(100)"`
	TimeAdd time.Time `json:"time_add" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	Aid     int       `json:"aid" xorm:"not null default 0 comment('管理员ID') INT(11)"`
	BlogId  int       `json:"blog_id" xorm:"not null default 0 comment('文章ID') INT(11)"`
}

//初始化
func NewBlogTag() *BlogTag {
	return new(BlogTag)
}
