package biz

import (
	"github.com/foxiswho/shop-go/service/order_service/biz/order_format"
	"github.com/foxiswho/shop-go/models/session_models"
	"github.com/foxiswho/shop-go/models/service_models/order_service_models"
	"github.com/foxiswho/shop-go/models"
	"fmt"
	"testing"
	"github.com/foxiswho/shop-go/service/order_service/biz/order_create"
)

func Test_Order(t *testing.T) {
	orderCreate_test(t)
}

func orderCreate_test(t *testing.T) {
	//用户数据
	user := &session_models.User{}
	user.Id = 1
	user.GroupId = 1
	user.TypePrice = 0
	user.Name = "测试数据"
	user.Username = "测试数据"
	user.Mobile = "18251188552"
	//
	goods := []*order_service_models.Goods{}
	price := &order_service_models.Goods{}
	price.Price = 100
	price.GoodsId = 1002
	price.Num = 3
	price.PriceId = 1002
	goods = append(goods, price)
	//
	consignee := &models.OrderConsignee{}
	consignee.Mobile = "18251188552"
	consignee.Consignee = "foxwho"
	consignee.ProvinceName = "江苏"
	consignee.CityName = "苏州"
	consignee.DistrictName = "园区"
	consignee.Address = "若水路388号"
	//
	order_format := order_format.NewOrderFormat()
	order_format.SetUser(user)
	order_format.SetGoods(goods)
	order_format.SetOrderConsignee(consignee)
	order_format.SetPayId(1)
	order_format.SetTypeId(1)
	err, order, order_goods := order_format.Process()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("NewOrderFormat order:", order)
	fmt.Println("NewOrderFormat order_goods:", order_goods)
	t.Log("NewOrderFormat order:", order)
	c := &order_create.OrderCreate{}
	c.SetOrder(order)
	c.SetOrderGoods(order_goods)
	err, orders := c.Process()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("NewOrderCreate :", orders)
}
