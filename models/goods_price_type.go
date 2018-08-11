package models

import (
	"time"
)

type GoodsPriceType struct {
	Id          int       `json:"id" xorm:"not null pk INT(11)"`
	Type        int       `json:"type" xorm:"not null default 0 comment('类型id') INT(11)"`
	Value       int       `json:"value" xorm:"not null default 0 comment('id') INT(11)"`
	GmtCreate   time.Time `json:"gmt_create" xorm:"default 'current_timestamp()' comment('添加时间') TIMESTAMP"`
	GmtModified time.Time `json:"gmt_modified" xorm:"default 'current_timestamp()' comment('更新时间') TIMESTAMP"`
	AidCreate   int       `json:"aid_create" xorm:"not null default 0 comment('添加人') INT(11)"`
	AidModified int       `json:"aid_modified" xorm:"not null default 0 comment('更新人') INT(1)"`

	//
	ExtData interface{} `json:"ExtData" xorm:"- <- ->"`
}

//初始化
func NewGoodsPriceType() *GoodsPriceType {
	return new(GoodsPriceType)
}
