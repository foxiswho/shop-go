package order_format

import (
	"github.com/foxiswho/shop-go/models/session_models"
	"github.com/foxiswho/shop-go/models/service_models/order_service_models"
	"github.com/foxiswho/shop-go/models"
	"github.com/foxiswho/shop-go/module/number"
)

//订单
type Order struct {
	payId          int     //支付方式
	typeId         int     //类别
	orderStatus    int     //订单状态
	orderSn        string  //自定义单号
	discount       float64 //优惠金额
	useWalletMoney float64 //使用钱包
	useCredit      int     //使用积分
	warehouseId    int     //仓库
	sid            int     //供应商ID
	isCustomPrice  bool    //是否使用自定义价格
	//
	user           session_models.User
	goods          []order_service_models.Goods //商品数据
	orderConsignee models.OrderConsignee
	ext            order_service_models.Ext
	order          *models.Order
	orderGoods     []*models.OrderGoods
}

func NewOrder() *Order {
	return new(Order)
}

//支付方式
func (c *Order) SetPayId(payId int) {
	c.payId = payId
}

//类别
func (c *Order) SetTypeId(typeId int) {
	c.typeId = typeId
}

//订单状态
func (c *Order) SetOrderStatus(orderStatus int) {
	c.orderStatus = orderStatus
}

//自定义单号
func (c *Order) SetOrderSn(orderSn string) {
	c.orderSn = orderSn
}

func (c *Order) SetDiscount(discount float64) {
	c.discount = discount
}
func (c *Order) SetUseWalletMoney(useWalletMoney float64) {
	c.useWalletMoney = useWalletMoney
}

func (c *Order) SetUseCredit(useCredit int) {
	c.useCredit = useCredit
}

func (c *Order) SetWarehouseId(warehouseId int) {
	c.warehouseId = warehouseId
}

func (c *Order) SetSid(sid int) {
	c.sid = sid
}

func (c *Order) SetIsCustomPrice(isCustomPrice bool) {
	c.isCustomPrice = isCustomPrice
}

func (c *Order) SetUser(user session_models.User) {
	c.user = user
}

func (c *Order) SetGoods(goods []order_service_models.Goods) {
	c.goods = goods
}

func (c *Order) SetOrderConsignee(orderConsignee models.OrderConsignee) {
	c.orderConsignee = orderConsignee
}

func (c *Order) SetExt(ext order_service_models.Ext) {
	c.ext = ext
}

func (c *Order) formatOrder() {
	order := new(models.Order)
	order.OrderNo = number.OrderNoMake()
	if len(c.orderSn) > 0 {
		order.OrderSn = c.orderSn
	}
	order.Uid = c.user.Id
	order.OrderStatus = c.orderStatus
	order.TypeId = c.typeId
	order.TypeIdAdmin = c.typeId
	order.AmountDiscount = c.discount
	order.UseWallet = c.useWalletMoney
	order.UseCredit = c.useCredit
	c.order = order
}
func (c *Order) formatGoods() {
	if c.goods != nil {
		for key, goods := range c.goods {
			order_goods := new(models.OrderGoods)
			order_goods.PriceId=goods.PriceId
			order_goods.GoodsId=goods.GoodsId
			c.orderGoods[key] = order_goods
		}
	}


}
