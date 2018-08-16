package order_create

import (
	"github.com/foxiswho/shop-go/models"
	"github.com/foxiswho/shop-go/module/db"
)

//订单
type OrderCreate struct {
	order             *models.Order
	orderGoods        []*models.OrderGoods
	orderGoodsStructs []*models.OrderGoodsStructure
}

func NewOrderCreate() *OrderCreate {
	return new(OrderCreate)
}

//
func (c *OrderCreate) SetOrder(order *models.Order) {
	c.order = order
}

func (c *OrderCreate) SetOrderGoods(orderGoods []*models.OrderGoods) {
	c.orderGoods = orderGoods
}

func (c *OrderCreate) Process() (error, *models.Order) {
	err := c.saveOrder()
	if err != nil {
		return err, nil
	}
	c.saveOrderGoods()

	return nil, c.order
}

func (c *OrderCreate) saveOrder() error {
	_, err := db.Db().Engine.Insert(c.order)
	if err != nil {
		return err
	}
	return nil
}

func (c *OrderCreate) saveOrderGoods() error {
	for _, item := range c.orderGoods {
		item.OrderId = c.order.Id
		_, err := db.Db().Engine.Insert(item)
		if err != nil {
			return err
		}
	}
	return nil
}
