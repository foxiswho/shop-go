package price

import (
	"github.com/foxiswho/shop-go/models"
	"github.com/foxiswho/shop-go/models/session_models"
	"github.com/foxiswho/shop-go/consts/goods"
	"strconv"
	"github.com/foxiswho/shop-go/dao/goods_dao"
)

type GoodsPrice struct {
	user             *session_models.User
	goodsPrice       *models.GoodsPrice
	prices           []*models.GoodsPrice
	pricesTypeFormat map[string]int
}

func NewGoodsPrice() *GoodsPrice {
	return new(GoodsPrice)
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
func (c *GoodsPrice) Process() float64 {
	var price_user float64 = 0
	//在URL中显示，并且不是默认类别时
	if 1 == c.goodsPrice.IsUrlShow && goods.Price_Type_Default != c.goodsPrice.PriceTypeSub && c.goodsPrice.IsDel == 0 {
		price_user = c.goodsPrice.PriceShop
	} else if goods.Price_Type_Default == c.goodsPrice.PriceTypeSub {
		//默认类别
		//有优惠价格存在时
		if len(c.prices) > 0 {
			//获取优惠属性价格
			c.getPricesTypeFormat()
			//匹配价格
			price_user = c.pricesMatch()
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
func (c *GoodsPrice) pricesMatch() (float64) {
	var price_user float64 = 0
	is_price_find := false
	for _, price := range c.prices {
		//如果是已删除价格则PASS
		if 1 == price.IsDel {
			continue
		}
		//如果是默认状态则PASS
		if goods.Price_Type_Default != price.PriceType {
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
	case goods.Price_Type_Sub_Default:
		price_user = price.PriceShop
		is_price_find = true
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

//获取优惠属性价格
func (c *GoodsPrice) getPricesTypeFormat() {
	c.pricesTypeFormat = make(map[string]int)
	if len(c.prices) > 0 {
		type_id := []int{}
		for _, item := range c.prices {
			//删除 PASS
			if 1 == item.IsDel {
				continue
			}
			//URL显示PASS
			if 1 == item.IsUrlShow {
				continue
			}
			//默认PASS
			if goods.Price_Type_Default == item.PriceTypeSub {
				continue
			}
			type_id = append(type_id, item.PriceTypeSub)
			//指定用户
			if goods.Price_Type_Sub_User == item.PriceTypeSub {
				if c.getPriceTypeByTypeId(item) {
					key := strconv.Itoa(item.Id) + "-" + strconv.Itoa(item.PriceTypeSub) + "-" + strconv.Itoa(c.user.Id)
					c.pricesTypeFormat[key] = 1
				}
				continue
			}
			//用户组
			if goods.Price_Type_Sub_Group == item.PriceTypeSub {
				if c.user.GroupId == item.GroupId {
					key := strconv.Itoa(item.Id) + "-" + strconv.Itoa(item.PriceTypeSub) + "-" + strconv.Itoa(c.user.GroupId)
					c.pricesTypeFormat[key] = 1
				}
				continue

			}
			//自定义类别
			if goods.Price_Type_Sub_Custom_Type == item.PriceTypeSub {
				if c.getPriceTypeByTypeId(item) {
					key := strconv.Itoa(item.Id) + "-" + strconv.Itoa(item.PriceTypeSub) + "-" + strconv.Itoa(c.user.TypePrice)
					c.pricesTypeFormat[key] = 1
				}
				continue
			}
		}
	}
}

//在指定价格 是否有指定类别
func (c *GoodsPrice) getPriceTypeByTypeId(price *models.GoodsPrice) bool {
	//TODO  后期改成从 缓存取值
	return goods_dao.GetPriceTypeByTypeId(c.user, price)
}
