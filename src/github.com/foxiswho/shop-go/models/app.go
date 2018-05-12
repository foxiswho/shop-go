package models

import (
	"time"
)

type App struct {
	AppId   int       `json:"app_id" xorm:"not null pk autoincr INT(11)"`
	TypeId  int       `json:"type_id" xorm:"not null default 0 comment('app_id,来源type表') unique INT(11)"`
	Name    string    `json:"name" xorm:"comment('名称') VARCHAR(100)"`
	Mark    string    `json:"mark" xorm:"comment('标志') CHAR(32)"`
	Setting string    `json:"setting" xorm:"comment('扩展参数') VARCHAR(5000)"`
	Remark  string    `json:"remark" xorm:"comment('备注') VARCHAR(255)"`
	TimeAdd time.Time `json:"time_add" xorm:"default 'CURRENT_TIMESTAMP' comment('添加时间') TIMESTAMP"`
	IsDel   int       `json:"is_del" xorm:"not null default 0 comment('是否删除0否1是') INT(11)"`
}

//初始化
func NewApp() *App {
	return new(App)
}
