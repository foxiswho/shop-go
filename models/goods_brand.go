package models

import (
	"time"
)

type GoodsBrand struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Name        string    `json:"name" xorm:"default 'NULL' comment('名称') VARCHAR(255)"`
	NameEn      string    `json:"name_en" xorm:"default 'NULL' comment('品牌英文名称或是汉语拼音') VARCHAR(255)"`
	Http        string    `json:"http" xorm:"default 'NULL' comment('品牌网站') VARCHAR(255)"`
	Phone       string    `json:"phone" xorm:"default 'NULL' comment('客服电话') VARCHAR(255)"`
	Content     string    `json:"content" xorm:"default 'NULL' comment('品牌介绍') TEXT"`
	Letter      string    `json:"letter" xorm:"default 'NULL' comment('品牌首字母') VARCHAR(255)"`
	Sort        int       `json:"sort" xorm:"not null default 0 comment('品牌排序') INT(10)"`
	Logo        string    `json:"logo" xorm:"default 'NULL' comment('品牌logo') VARCHAR(255)"`
	ParentId    int       `json:"parent_id" xorm:"not null default 0 index INT(10)"`
	ArrParentId string    `json:"arr_parent_id" xorm:"default 'NULL' comment('所有父栏目ID') VARCHAR(255)"`
	IsChild     int       `json:"is_child" xorm:"not null default 0 comment('是否有子栏目') TINYINT(1)"`
	ArrChildId  string    `json:"arr_child_id" xorm:"default 'NULL' comment('所有子栏目ID') TEXT"`
	IsDel       int       `json:"is_del" xorm:"not null default 0 comment('是否删除1是0否') index TINYINT(1)"`
	GmtCreate   time.Time `json:"gmt_create" xorm:"not null default 'current_timestamp()' comment('添加时间') TIMESTAMP"`
	GmtModified time.Time `json:"gmt_modified" xorm:"default 'current_timestamp()' comment('更新时间') TIMESTAMP"`

	//
	ExtData interface{} `json:"ExtData" xorm:"- <- ->"`
}

//初始化
func NewGoodsBrand() *GoodsBrand {
	return new(GoodsBrand)
}
