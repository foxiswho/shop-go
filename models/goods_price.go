package models

import (
	"time"
)

type GoodsPrice struct {
	Id              int       `json:"id" xorm:"not null pk autoincr INT(11)"`
	GoodsId         int       `json:"goods_id" xorm:"not null default 0 comment('商品ID') INT(11)"`
	Name            string    `json:"name" xorm:"default 'NULL' comment('商品名称') VARCHAR(200)"`
	PriceMarket     int64     `json:"price_market" xorm:"not null default 0 comment('市场价') BIGINT(20)"`
	PriceShop       int64     `json:"price_shop" xorm:"not null default 0 comment('商城价') BIGINT(20)"`
	PriceType       int       `json:"price_type" xorm:"not null default 0 comment('价格类别') INT(11)"`
	IsFreeDelivery  int       `json:"is_free_delivery" xorm:"not null default 0 comment('是否包邮1是0否') TINYINT(1)"`
	TimeStart       time.Time `json:"time_start" xorm:"default 'NULL' comment('开始时间') DATETIME"`
	TimeEnd         time.Time `json:"time_end" xorm:"default 'NULL' comment('结束时间') DATETIME"`
	MinFreeDelivery int       `json:"min_free_delivery" xorm:"not null default 1 comment('最小包邮数量') INT(10)"`
	NumMax          int       `json:"num_max" xorm:"not null default 9999 comment('最大可一次购买数量') INT(11)"`
	NumLeast        int       `json:"num_least" xorm:"not null default 1 comment('最少购买数量') INT(11)"`
	IsGroupPrice    int       `json:"is_group_price" xorm:"not null default 1 comment('是否使用用户组价格') TINYINT(1)"`
}

//初始化
func NewGoodsPrice() *GoodsPrice {
	return new(GoodsPrice)
}
