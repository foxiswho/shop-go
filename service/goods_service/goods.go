package goods_service

import (
	"github.com/foxiswho/shop-go/models/session_models"
	"github.com/foxiswho/shop-go/models"
	"github.com/foxiswho/shop-go/service/goods_service/price"
)

//获取价格 单个商品价格
func Price(user *session_models.User, goodsPrice *models.GoodsPrice, prices []*models.GoodsPrice) float64 {
	c := price.NewGoodsPrice()
	c.SetUser(user)
	c.SetGoodsPrice(goodsPrice)
	c.SetPrices(prices)
	return c.Process()
}

//获取价格 单个商品价格
func PriceByGoodsPrice(user *session_models.User, goodsPrice *models.GoodsPrice) float64 {
	return Price(user, goodsPrice, price.GetPricesByPrice(goodsPrice))
}
