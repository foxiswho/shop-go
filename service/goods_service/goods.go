package goods_service

import (
	"github.com/foxiswho/shop-go/models"
	"github.com/foxiswho/shop-go/models/session_models"
	"github.com/foxiswho/shop-go/consts/goods"
	"strconv"
)

type GoodsPrice struct {
	user             *session_models.User
	goodsPrice       *models.GoodsPrice
	prices           []*models.GoodsPrice
	pricesTypeFormat map[string]int
}

func (c *GoodsPrice) SetUser(user *session_models.User) {
	c.user = user
}

func (c *GoodsPrice) SetGoodsPrice(price *models.GoodsPrice) {
	c.goodsPrice = price
}

func (c *GoodsPrice) SetPrices(price []*models.GoodsPrice) {
	c.prices = price
}
func (c *GoodsPrice) processPrice() float64 {
	var price_user float64 = 0
	//在URL中显示，并且不是默认类别时
	if 1 == c.goodsPrice.IsUrlShow && goods.Price_Type_Default != c.goodsPrice.PriceTypeSub && c.goodsPrice.IsDel == 0 {
		price_user = c.goodsPrice.PriceShop
	} else if goods.Price_Type_Default == c.goodsPrice.PriceTypeSub {
		//默认类别
		//有优惠价格存在时
		if len(c.prices) > 0 {
			is_price_find := false
			for _, price := range c.prices {
				//如果是已删除价格则PASS
				if price.IsDel {
					continue
				}
				//如果是默认状态则PASS
				if goods.Price_Type_Default != price.PriceTypeSub {
					continue
				}
				//指定用户
				if goods.Price_Type_Sub_User == price.PriceTypeSub {
					//是否存在指定用户价格
					if c.isUserPriceTypeExists(price) {
						price_user = price.PriceShop
						is_price_find = true
						break
					}
				} else if goods.Price_Type_Sub_Group == price.PriceTypeSub {
					//指定用户组
					//是否存在指定用户组价格
					if c.isGroupPriceTypeExists(price) {
						price_user = price.PriceShop
						is_price_find = true
						break
					}
				} else if goods.Price_Type_Sub_Group == price.PriceTypeSub {
					//自定义类别
					//是否存在自定义类别价格
					if c.isCustomPriceTypeExists(price) {
						price_user = price.PriceShop
						is_price_find = true
						break
					}
				}
			}
			//没有匹配到价格，那么使用当前价格
			if false == is_price_find {
				//如果是删除的
				if c.goodsPrice.IsDel == 1 {
					price_user = 0
				} else {
					price_user = c.goodsPrice.PriceShop
				}
			}
		} else {
			//如果是删除的
			if c.goodsPrice.IsDel == 1 {
				price_user = 0
			} else {
				price_user = c.goodsPrice.PriceShop
			}
		}
	}
	return price_user
}

//优惠价格匹配
func (c *GoodsPrice) pricesMatch() float64 {
	var price_user float64 = 0
	is_price_find := false
	for _, price := range c.prices {
		//如果是已删除价格则PASS
		if price.IsDel {
			continue
		}
		//如果是默认状态则PASS
		if goods.Price_Type_Default != price.PriceTypeSub {
			continue
		}
		is_price_find, price_user = c.priceMatching(price)
	}
	//没有匹配到价格，那么使用当前价格
	if false == is_price_find {
		//如果是删除的
		if c.goodsPrice.IsDel == 1 {
			price_user = 0
		} else {
			price_user = c.goodsPrice.PriceShop
		}
	}
	return price_user
}

//实际匹配
func (c *GoodsPrice) priceMatching(price *models.GoodsPrice) (bool, float64) {
	var price_user float64 = 0
	is_price_find := false
	switch price.PriceTypeSub {
	case goods.Price_Type_Sub_User:
		//是否存在指定用户价格
		if c.isUserPriceTypeExists(price) {
			price_user = price.PriceShop
			is_price_find = true
		}
	case goods.Price_Type_Sub_Group:
		//指定用户组
		//是否存在指定用户组价格
		if c.isGroupPriceTypeExists(price) {
			price_user = price.PriceShop
			is_price_find = true
		}
	case goods.Price_Type_Sub_Custom_Type:
		//自定义类别
		//是否存在自定义类别价格
		if c.isCustomPriceTypeExists(price) {
			price_user = price.PriceShop
			is_price_find = true
		}
	}
	return is_price_find, price_user
}

//指定用户价格是否存在
func (c *GoodsPrice) isUserPriceTypeExists(price *models.GoodsPrice) bool {
	//键名
	key := strconv.Itoa(price.Id) + "-" + strconv.Itoa(price.PriceTypeSub) + "-" + strconv.Itoa(c.user.Id)
	//查找是否存在
	if _, ok := c.pricesTypeFormat[key]; ok {
		return true
	}
	return false
}

//指定用户组价格是否存在
func (c *GoodsPrice) isGroupPriceTypeExists(price *models.GoodsPrice) bool {
	//键名
	key := strconv.Itoa(price.Id) + "-" + strconv.Itoa(price.PriceTypeSub) + "-" + strconv.Itoa(c.user.GroupId)
	//查找是否存在
	if _, ok := c.pricesTypeFormat[key]; ok {
		return true
	}
	return false
}

//自定义类别价格是否存在
func (c *GoodsPrice) isCustomPriceTypeExists(price *models.GoodsPrice) bool {
	//键名
	key := strconv.Itoa(price.Id) + "-" + strconv.Itoa(price.PriceTypeSub) + "-" + strconv.Itoa(c.user.TypePrice)
	//查找是否存在
	if _, ok := c.pricesTypeFormat[key]; ok {
		return true
	}
	return false
}

func Price(models_price models.GoodsPrice, user session_models.User, prices []*models.GoodsPrice) float64 {
	goods_price := models_price.PriceShop
	if len(prices) > 0 {
		for _, price := range prices {
			//如果是已删除价格则PASS
			if price.IsDel {
				continue
			}

		}
	} else {
		if models_price.IsDel {
			return 0
		}
	}
	return goods_price
}
func priceSwitch(price models.GoodsPrice) {
	if goods.Price_Type_Sub_User == price.PriceTypeSub {

	}
}
