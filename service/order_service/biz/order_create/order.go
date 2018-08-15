package order_create

import (
	"github.com/foxiswho/shop-go/models"
	"github.com/foxiswho/shop-go/module/db"
)

//订单
type Order struct {
	order             *models.Order
	orderGoods        []*models.OrderGoods
	orderGoodsStructs []*models.OrderGoodsStructure
}

func NewOrder() *Order {
	return new(Order)
}

//
func (c *Order) SetOrder(order *models.Order) {
	c.order = order
}

func (c *Order) SetOrderGoods(orderGoods []*models.OrderGoods) {
	c.orderGoods = orderGoods
}

func (c *Order) process() (error, *models.Order) {
	err := c.saveOrder()
	if err != nil {
		return err, nil
	}
	c.saveOrderGoods()

	return nil, c.order
}

func (c *Order) saveOrder() error {
	_, err := db.Db().Engine.Insert(c.order)
	if err != nil {
		return err
	}
	return nil
}

func (c *Order) saveOrderGoods() error {
	for _, item := range c.orderGoods {
		item.OrderId = c.order.Id
		_, err := db.Db().Engine.Insert(item)
		if err != nil {
			return err
		}
	}
	return nil
}
