package models

import (
	"time"
)

type GoodsPrice struct {
	Id              int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	GoodsId         int       `json:"goods_id" xorm:"not null default 0 comment('商品ID') INT(11)"`
	Name            string    `json:"name" xorm:"default 'NULL' comment('商品名称') VARCHAR(200)"`
	PriceMarket     float64   `json:"price_market" xorm:"not null default 0.0000 comment('市场价') DECIMAL(28,4)"`
	PriceShop       float64   `json:"price_shop" xorm:"not null default 0.0000 comment('商城价') DECIMAL(28,4)"`
	PriceType       int       `json:"price_type" xorm:"not null default 0 comment('价格类别') INT(11)"`
	PriceTypeSub    int       `json:"price_type_sub" xorm:"not null default 0 comment('价格类别子类别') INT(11)"`
	TimeStart       time.Time `json:"time_start" xorm:"default 'NULL' comment('开始时间') DATETIME"`
	TimeEnd         time.Time `json:"time_end" xorm:"default 'NULL' comment('结束时间') DATETIME"`
	NumMax          int       `json:"num_max" xorm:"not null default 9999 comment('最大可一次购买数量') INT(11)"`
	NumMin          int       `json:"num_min" xorm:"not null default 1 comment('最少购买数量') INT(11)"`
	MinFreeDelivery int       `json:"min_free_delivery" xorm:"not null default 1 comment('最小包邮数量') INT(10)"`
	IsFreeDelivery  int       `json:"is_free_delivery" xorm:"not null default 0 comment('是否包邮1是0否') TINYINT(1)"`
	IsGroupPrice    int       `json:"is_group_price" xorm:"not null default 1 comment('是否使用用户组价格') TINYINT(1)"`
	IsFreeTax       int       `json:"is_free_tax" xorm:"not null default 0 comment('是否包税') TINYINT(1)"`
	IsDel           int       `json:"is_del" xorm:"not null default 0 comment('是否删除') TINYINT(1)"`
	GroupId         int       `json:"group_id" xorm:"not null comment('指定用户组') INT(11)"`
	TypePrice       int       `json:"type_price" xorm:"not null comment('指定价格类别') INT(11)"`
	GmtCreate       time.Time `json:"gmt_create" xorm:"default 'current_timestamp()' comment('添加时间') TIMESTAMP"`
	GmtModified     time.Time `json:"gmt_modified" xorm:"default 'current_timestamp()' comment('更新时间') TIMESTAMP"`
	AidCreate       int       `json:"aid_create" xorm:"not null default 0 comment('添加人') INT(11)"`
	AidModified     int       `json:"aid_modified" xorm:"not null default 0 comment('更新人') INT(1)"`

	//
	ExtData interface{} `json:"ExtData" xorm:"- <- ->"`
}

//初始化
func NewGoodsPrice() *GoodsPrice {
	return new(GoodsPrice)
}
