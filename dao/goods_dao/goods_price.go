package goods_dao

import (
	"github.com/foxiswho/shop-go/models"
	"fmt"
	"github.com/foxiswho/shop-go/module/db"
	"github.com/foxiswho/shop-go/models/session_models"
)

//取出 当前价格所属商品的 所有优惠价格
func GetPricesByPrice(price *models.GoodsPrice) []*models.GoodsPrice {
	//TODO  后期 优化
	prices := make([]*models.GoodsPrice, 0)
	err := db.Db().Engine.Where("goods_id =? and is_url_show=0 and is_del=0", price.GoodsId).OrderBy("price_type_sub ASC").Find(&prices)
	if err != nil {
		fmt.Println(err)
	}
	return prices
}

//在指定价格 是否有指定类别
func GetPriceTypeByTypeId(user *session_models.User, price *models.GoodsPrice) bool {
	//TODO  后期改成从 缓存取值
	p := &models.GoodsPriceType{}
	_, err := db.Db().Engine.Where("price_id=? and type=? and value=?", price.Id, price.PriceTypeSub, user.GroupId).Get(&p)
	if err != nil {
		return false
	}
	if p.Id > 0 {
		return true
	}
	return false
}
