package order_format

import (
	"github.com/foxiswho/shop-go/models/session_models"
	"github.com/foxiswho/shop-go/models/service_models/order_service_models"
	"github.com/foxiswho/shop-go/models"
	"github.com/foxiswho/shop-go/module/number"
	"github.com/foxiswho/shop-go/dao/crud"
	"github.com/foxiswho/shop-go/util"
	"github.com/foxiswho/shop-go/service/goods_service"
)

//订单
type OrderFormat struct {
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
	user           *session_models.User
	goods          []*order_service_models.Goods //商品数据
	orderConsignee *models.OrderConsignee
	ext            *order_service_models.Ext
	//
	order             *models.Order
	orderGoods        []*models.OrderGoods
	orderGoodsStructs []*models.OrderGoodsStructure
	//
	goodsStructs map[int]*models.GoodsStruct //商品价格数据
}

func NewOrderFormat() *OrderFormat {
	return new(OrderFormat)
}

//支付方式
func (c *OrderFormat) SetPayId(payId int) {
	c.payId = payId
}

//类别
func (c *OrderFormat) SetTypeId(typeId int) {
	c.typeId = typeId
}

//订单状态
func (c *OrderFormat) SetOrderStatus(orderStatus int) {
	c.orderStatus = orderStatus
}

//自定义单号
func (c *OrderFormat) SetOrderSn(orderSn string) {
	c.orderSn = orderSn
}

func (c *OrderFormat) SetDiscount(discount float64) {
	c.discount = discount
}
func (c *OrderFormat) SetUseWalletMoney(useWalletMoney float64) {
	c.useWalletMoney = useWalletMoney
}

func (c *OrderFormat) SetUseCredit(useCredit int) {
	c.useCredit = useCredit
}

func (c *OrderFormat) SetWarehouseId(warehouseId int) {
	c.warehouseId = warehouseId
}

func (c *OrderFormat) SetSid(sid int) {
	c.sid = sid
}

func (c *OrderFormat) SetIsCustomPrice(isCustomPrice bool) {
	c.isCustomPrice = isCustomPrice
}

func (c *OrderFormat) SetUser(user *session_models.User) {
	c.user = user
}

func (c *OrderFormat) SetGoods(goods []*order_service_models.Goods) {
	c.goods = goods
}

func (c *OrderFormat) SetOrderConsignee(orderConsignee *models.OrderConsignee) {
	c.orderConsignee = orderConsignee
}

func (c *OrderFormat) SetExt(ext *order_service_models.Ext) {
	c.ext = ext
}
func (c *OrderFormat) Process() (error, *models.Order, []*models.OrderGoods) {
	//订单数据
	c.formatOrder()
	//整合商品价格数据
	err := c.formatGoodsPrice()
	if err != nil {
		return err, nil, nil
	}
	c.formatOrderGoods()
	return nil, c.order, c.orderGoods
}

//订单数据格式化
func (c *OrderFormat) formatOrder() {
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
func (c *OrderFormat) formatGoodsPrice() error {
	if c.goods != nil && len(c.goods) > 0 {
		ids := []int{}
		ids_price := []int{}
		for _, goods := range c.goods {
			ids = append(ids, goods.GoodsId)
			ids_price = append(ids_price, goods.PriceId)
		}
		//获取商品数据，并按商品ID索引
		data_goods, err := crud.NewGoodsCrud().GetByIdsIndex(ids)
		if err != nil {
			return err
		}
		//获取价格数据
		data_price, err := crud.NewGoodsPriceCrud().GetByIds(ids_price)
		if err != nil {
			return err
		}
		warehouse_id := 0
		sid := 0
		//整合
		goods_struct := make(map[int]*models.GoodsStruct)
		for _, item := range data_price {
			one := models.NewGoodsStruct()
			one.Goods = data_goods[item.Id]
			one.GoodsPrice = &item
			goods_struct[item.Id] = one
			if warehouse_id == 0 {
				warehouse_id = one.Goods.WarehouseId
			} else {
				if warehouse_id != one.Goods.WarehouseId {
					return util.NewError("商品必须是 同一仓库")
				}
			}
			if sid == 0 {
				sid = one.Goods.Sid
			} else {
				if sid != one.Goods.Sid {
					return util.NewError("商品必须是 同一供应商")
				}
			}
		}
		c.goodsStructs = goods_struct
		c.order.Sid = sid
		c.order.WarehouseId = warehouse_id
		return nil
	}
	return util.NewError("商品数据错误")
}
func (c *OrderFormat) formatOrderGoods() {
	c.orderGoods = []*models.OrderGoods{}
	for _, item := range c.goods {
		goods := c.goodsStructs[item.PriceId]
		one := &models.OrderGoods{}
		one.PriceId = item.PriceId
		//使用自定义价格
		if c.isCustomPrice {
			one.Price = item.Price
		} else {
			//从商品上获取最新价格
			one.Price = goods_service.PriceByGoodsPrice(c.user, goods.GoodsPrice)
		}
		one.PriceShop = goods.GoodsPrice.PriceShop
		one.GoodsId = goods.Goods.Id
		one.ProductId = goods.Goods.ProductId
		one.Num = item.Num
		one.NumUnit = goods.Goods.NumUnit
		one.NumTotal = one.Num * one.NumUnit
		one.Amount = one.Price * float64(one.Num)
		one.CostPrice = goods.GoodsPrice.PriceShop
		one.CostAmount = one.CostPrice * float64(one.Num)
		one.Title = goods.GoodsPrice.Name
		one.Number = goods.Goods.Number
		one.WarehouseId = goods.Goods.WarehouseId
		one.Sid = goods.Goods.Sid
		one.MarkId = goods.Goods.MarkId
		c.orderGoods = append(c.orderGoods, one)
		//订单价格
		c.order.AmountGoods += one.Amount
	}
}
