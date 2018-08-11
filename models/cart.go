package models

import (
	"time"
)

type Cart struct {
	Id          int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	Uid         int       `json:"uid" xorm:"not null default 0 comment('用户ID') index INT(10)"`
	GoodsId     int       `json:"goods_id" xorm:"not null default 0 comment('商品ID') INT(10)"`
	ProductId   int       `json:"product_id" xorm:"not null default 0 comment('商品信息id') INT(10)"`
	Num         int       `json:"num" xorm:"not null default 0 comment('数量') INT(10)"`
	Price       float64   `json:"price" xorm:"not null default 0.0000 comment('单价') DECIMAL(28,4)"`
	Amount      float64   `json:"amount" xorm:"not null default 0.0000 comment('合计总价') DECIMAL(28,4)"`
	WarehouseId int       `json:"warehouse_id" xorm:"not null default 0 comment('仓库ID') INT(10)"`
	Sid         int       `json:"sid" xorm:"not null default 0 comment('供货商ID') INT(10)"`
	TypeId      int       `json:"type_id" xorm:"not null default 1 comment('类别:普通') index INT(11)"`
	GmtCreate   time.Time `json:"gmt_create" xorm:"not null default 'current_timestamp()' comment('添加时间') TIMESTAMP"`
	GmtModified time.Time `json:"gmt_modified" xorm:"default 'current_timestamp()' comment('更新时间') TIMESTAMP"`

	//
	ExtData interface{} `json:"ExtData" xorm:"- <- ->"`
}

//初始化
func NewCart() *Cart {
	return new(Cart)
}
