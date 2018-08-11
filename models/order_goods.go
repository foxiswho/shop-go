package models

type OrderGoods struct {
	Id          int     `json:"id" xorm:"not null pk autoincr INT(10)"`
	OrderId     int     `json:"order_id" xorm:"not null default 0 comment('订单ID') index INT(10)"`
	PriceId     int     `json:"price_id" xorm:"not null default 0 comment('价格id') INT(11)"`
	GoodsId     int     `json:"goods_id" xorm:"not null default 0 comment('商品ID') INT(10)"`
	ProductId   int     `json:"product_id" xorm:"not null default 0 comment('商品信息id') INT(10)"`
	Title       string  `json:"title" xorm:"default 'NULL' comment('商品名称') VARCHAR(200)"`
	Num         int     `json:"num" xorm:"not null default 0 comment('数量') INT(10)"`
	Number      string  `json:"number" xorm:"default 'NULL' comment('商品编号') CHAR(100)"`
	Price       float64 `json:"price" xorm:"not null default 0.0000 comment('单价') DECIMAL(28,4)"`
	NumUnit     int     `json:"num_unit" xorm:"not null default 1 comment('每个单位内多少个，每盒几罐') INT(11)"`
	NumTotal    int     `json:"num_total" xorm:"not null default 0 comment('总数量 = 罐数x页面数量') INT(11)"`
	Amount      float64 `json:"amount" xorm:"not null default 0.0000 comment('合计总价') DECIMAL(28,4)"`
	Freight     float64 `json:"freight" xorm:"not null default 0.0000 comment('运费') DECIMAL(28,4)"`
	WarehouseId int     `json:"warehouse_id" xorm:"not null default 0 comment('仓库ID') INT(10)"`
	Sid         int     `json:"sid" xorm:"not null default 1 comment('商家ID') INT(10)"`
	SalesFee    float64 `json:"sales_fee" xorm:"not null default 0.0000 comment('消费税费') DECIMAL(28,4)"`
	VatFee      float64 `json:"vat_fee" xorm:"not null default 0.0000 comment('增值税费') DECIMAL(28,4)"`
	PriceTax    float64 `json:"price_tax" xorm:"not null default 0.0000 comment('总税费') DECIMAL(28,4)"`
	Remark      string  `json:"remark" xorm:"default 'NULL' comment('备注') TEXT"`
	PriceShop   float64 `json:"price_shop" xorm:"not null default 0.0000 comment('商城价') DECIMAL(28,4)"`
	CostPrice   float64 `json:"cost_price" xorm:"not null default 0.0000 comment('成本单价') DECIMAL(28,4)"`
	CostAmount  float64 `json:"cost_amount" xorm:"not null default 0.0000 comment('成本金额') DECIMAL(28,4)"`
	MarkId      int     `json:"mark_id" xorm:"not null default 0 comment('商品标志ID') INT(11)"`

	//
	ExtData interface{} `json:"ExtData" xorm:"- <- ->"`
}

//初始化
func NewOrderGoods() *OrderGoods {
	return new(OrderGoods)
}
