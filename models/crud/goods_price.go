
package crud

import (
	"fmt"
	"github.com/foxiswho/shop-go/models"
	"github.com/foxiswho/shop-go/module/db"
	"github.com/foxiswho/shop-go/util"
)

type GoodsPriceCrud struct {

}

func NewGoodsPriceCrud() *GoodsPriceCrud {
	return new(GoodsPriceCrud)
}

//初始化列表
func goods_priceNewMakeDataArr() []models.GoodsPrice {
	return make([]models.GoodsPrice, 0)
}

//列表查询
func (s *GoodsPriceCrud) GetAll(where []*db.QueryCondition, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
	m := models.NewGoodsPrice()
	session := db.Filter(where)
	count, err := session.Count(m)
	if err != nil {
		fmt.Println(err)
		return nil, util.NewError(err.Error())
	}
	Query := db.Pagination(int(count), page, limit)
	if count == 0 {
		return Query, nil
	}

	session = db.Filter(where)
	if orderBy != "" {
		session.OrderBy(orderBy)
	}
	session.Limit(limit, Query.Offset)
	if len(fields) == 0 {
		session.AllCols()
	}
	data := goods_priceNewMakeDataArr()
	err = session.Find(&data)
	if err != nil {
		fmt.Println(err)
		return nil, util.NewError(err.Error())
	}
	Query.Data = make([]interface{}, len(data))
	for y, x := range data {
		Query.Data[y] = x
	}
	return Query, nil
}


// 获取 单条记录
func (s *GoodsPriceCrud) GetById(id int) (*models.GoodsPrice, error) {
    m:=new(models.GoodsPrice)
	m.Id = id
	ok, err := db.Db().Engine.Get(m)
    if err != nil {
        return nil, err
    }
    if !ok{
        return nil,util.NewError("数据不存在:"+err.Error())
    }
    return m, nil
}


// 获取 多条记录
func (s *GoodsPriceCrud) GetByIds(id []int) ([]models.GoodsPrice, error) {
	data := goods_priceNewMakeDataArr()
	err := db.Db().Engine.In("id", id).Find(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// 获取 多条记录
func (s *GoodsPriceCrud) GetByIdsIndex(id []int) (map[int]*models.GoodsPrice, error) {
	data := goods_priceNewMakeDataArr()
	err := db.Db().Engine.In("id", id).Find(&data)
	if err != nil {
		return nil, err
	}
	data_index := make(map[int]*models.GoodsPrice)
	for _, val := range data {
		data_index[val.Id] = &val
	}
	return data_index, nil
}


// 删除 单条记录
func (s *GoodsPriceCrud) Delete(id int) (int64, error) {
	m:=new(models.GoodsPrice)
	m.Id = id
	num, err := db.Db().Engine.Delete(m)
	if err == nil {
		return num, nil
	}
	return num, err
}