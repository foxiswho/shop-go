package order_service_models

import (
	"github.com/foxiswho/shop-go/models"
	"github.com/foxiswho/shop-go/models/session_models"
)

//优惠券
type Coupon struct {
	CouponGetId int    //优惠券ID
	CouponCode  string //优惠码
}

//商品数据
type Goods struct {
	PriceId int //价格ID
	GoodsId int //商品ID
	Price   int //单价
}

//扩展数据
type Ext struct {
	ExpressId int
	ExpressNo string
}

//订单
type Order struct {
	PayId          int    //支付方式
	TypeId         int    //类别
	OrderStatus    string //订单状态
	OrderSn        string //自定义单号
	Discount       int    //优惠金额
	UseWalletMoney int    //使用钱包
	UseCredit      int    //使用积分
	WarehouseId    int    //仓库
	Sid            int    //供应商ID
	IsCustomPrice  bool   //是否使用自定义价格
	//
	User           session_models.User
	GoodsData      []Goods //商品数据
	OrderConsignee models.OrderConsignee
	Ext            Ext
}
