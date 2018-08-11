package models

import (
	"time"
)

type Goods struct {
	Id                 int       `json:"id" xorm:"not null pk autoincr INT(10)"`
	ProductId          int       `json:"product_id" xorm:"not null default 0 comment('商品信息ID') INT(10)"`
	WarehouseId        int       `json:"warehouse_id" xorm:"not null default 0 comment('仓库ID') INT(10)"`
	Sid                int       `json:"sid" xorm:"not null default 0 comment('供应商ID') index INT(11)"`
	Status             int       `json:"status" xorm:"not null default 0 comment('状态0未审核99已审核') index TINYINT(1)"`
	IsDel              int       `json:"is_del" xorm:"not null default 0 comment('是否删除1是0否') index TINYINT(1)"`
	IsOpen             int       `json:"is_open" xorm:"not null default 0 comment('是否上架1是0否') index TINYINT(1)"`
	Aid                int       `json:"aid" xorm:"not null default 0 comment('管理员（发布人）ID') INT(10)"`
	CatId              int       `json:"cat_id" xorm:"not null default 1 comment('栏目id') index INT(10)"`
	BrandId            int       `json:"brand_id" xorm:"not null default 0 comment('品牌') index INT(10)"`
	Title              string    `json:"title" xorm:"default 'NULL' comment('标题') VARCHAR(100)"`
	Model              string    `json:"model" xorm:"default 'NULL' comment('规格') VARCHAR(100)"`
	Number             string    `json:"number" xorm:"default 'NULL' comment('商品编号') CHAR(100)"`
	Thumb              string    `json:"thumb" xorm:"default 'NULL' comment('缩略图') VARCHAR(255)"`
	OriginalImg        string    `json:"original_img" xorm:"default 'NULL' comment('原始图') VARCHAR(255)"`
	Sort               int       `json:"sort" xorm:"not null default 0 comment('排序') index INT(10)"`
	PriceBase          float64   `json:"price_base" xorm:"not null default 0.0000 comment('底价') DECIMAL(28,4)"`
	PricePlantformCost float64   `json:"price_plantform_cost" xorm:"not null default 0.0000 comment('平台成本') DECIMAL(28,4)"`
	AttrTypeId         int       `json:"attr_type_id" xorm:"not null default 0 comment('属性类别ID') INT(10)"`
	NumUnit            int       `json:"num_unit" xorm:"not null default 1 comment('每个单位内多少个，每盒几罐') INT(11)"`
	TypeStock          int       `json:"type_stock" xorm:"not null default 0 comment('是否仓库库存') INT(10)"`
	TypeId             int       `json:"type_id" xorm:"not null default 10001 comment('类别类目') INT(11)"`
	TypeDelivery       int       `json:"type_delivery" xorm:"not null default 10201 comment('商品发货属性类别') INT(11)"`
	Mark               string    `json:"mark" xorm:"not null default '''' comment('标志:产品-仓库-供应商') index CHAR(32)"`
	MarkId             int       `json:"mark_id" xorm:"not null default 10001 comment('标志类别') index INT(11)"`
	Description        string    `json:"description" xorm:"default 'NULL' comment('描述') VARCHAR(255)"`
	GmtCreate          time.Time `json:"gmt_create" xorm:"default 'current_timestamp()' comment('添加时间') TIMESTAMP"`
	GmtModified        time.Time `json:"gmt_modified" xorm:"default 'current_timestamp()' comment('更新时间') TIMESTAMP"`
	AidCreate          int       `json:"aid_create" xorm:"not null default 0 comment('添加人') index INT(11)"`
	AidModified        int       `json:"aid_modified" xorm:"not null default 0 comment('更新人') INT(1)"`

	//
	ExtData interface{} `json:"ExtData" xorm:"- <- ->"`
}

//初始化
func NewGoods() *Goods {
	return new(Goods)
}
