package price
//
//import (
//	"github.com/foxiswho/shop-go/models"
//	"github.com/foxiswho/shop-go/consts/goods"
//	"github.com/foxiswho/shop-go/module/db"
//	"fmt"
//	"github.com/foxiswho/shop-go/models/session_models"
//)
////暂时废弃
//type GoodsPriceData struct {
//	user  *session_models.User
//	price *models.GoodsPrice
//}
//
//func (c *GoodsPriceData) Process() {
//	//默认 类别
//	if goods.Price_Type_Default == c.price.PriceType {
//		prices := make([]models.GoodsPrice, 0)
//		err := db.Db().Engine.Where("price_type<>? and is_url_show=0 and is_del=0", goods.Price_Type_Default).OrderBy("price_type_sub ASC").Find(&prices)
//		if err != nil {
//			fmt.Errorf(err)
//		}
//		//有优惠价格
//		if len(prices) > 0 {
//			is_find_price := false
//			for _, item := range prices {
//				//指定用户
//				if false == is_find_price && goods.Price_Type_Sub_User == item.PriceTypeSub {
//					if c.getPriceTypeByUid(item) {
//						is_find_price = true
//						return item
//					}
//					continue
//				}
//				//用户组
//				if false == is_find_price && goods.Price_Type_Sub_Group == item.PriceTypeSub {
//					if c.user.GroupId == item.GroupId {
//						is_find_price = true
//						return item
//					}
//					continue
//
//				}
//				//自定义类别
//				if false == is_find_price && goods.Price_Type_Sub_Custom_Type == item.PriceTypeSub {
//					if c.getPriceTypeByCustomTypeId(item) {
//						is_find_price = true
//						return item
//					}
//					continue
//				}
//
//			}
//			//没有找到 优惠价格，直接返回当前价格
//			if false == is_find_price {
//				return c.price
//			}
//		} else {
//			//没有优惠价格，直接返回当前价格
//			return c.price
//		}
//	} else if goods.Price_Type_Discount == c.price.PriceType && 1 == c.price.IsUrlShow {
//		// 显示连接
//		//TODO 促销类别没有去判断，后续再添加此地方，这里直接使用该价格
//	} else {
//		//空
//	}
//}
//
////在指定价格 是否有指定用户
//func (c *GoodsPriceData) getPriceTypeByUid(price *models.GoodsPrice) bool {
//	//TODO  后期改成从 缓存取值
//	p := &models.GoodsPriceType{}
//	_, err := db.Db().Engine.Where("price_id=? and type=? and value=?", price.Id, price.PriceTypeSub, c.user.Id).Get(&p)
//	if err != nil {
//		return false
//	}
//	if p.Id > 0 {
//		return true
//	}
//	return false
//}
//
////在指定价格 是否有指定类别
//func (c *GoodsPriceData) getPriceTypeByCustomTypeId(price *models.GoodsPrice) bool {
//	//TODO  后期改成从 缓存取值
//	p := &models.GoodsPriceType{}
//	_, err := db.Db().Engine.Where("price_id=? and type=? and value=?", price.Id, price.PriceTypeSub, c.user.GroupId).Get(&p)
//	if err != nil {
//		return false
//	}
//	if p.Id > 0 {
//		return true
//	}
//	return false
//}
