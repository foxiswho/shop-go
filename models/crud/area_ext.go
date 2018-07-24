
package crud

import (
	"fmt"
	"github.com/foxiswho/shop-go/models"
	"github.com/foxiswho/shop-go/module/db"
	"github.com/foxiswho/shop-go/util"
)

type AreaExtCrud struct {

}

func NewAreaExtCrud() *AreaExtCrud {
	return new(AreaExtCrud)
}

//初始化列表
func area_extNewMakeDataArr() []models.AreaExt {
	return make([]models.AreaExt, 0)
}

//列表查询
func (s *AreaExtCrud) GetAll(where []*db.QueryCondition, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
	m := models.NewAreaExt()
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
	data := area_extNewMakeDataArr()
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
func (s *AreaExtCrud) GetById(id int) (*models.AreaExt, error) {
    m:=new(models.AreaExt)
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

// 删除 单条记录
func (s *AreaExtCrud) Delete(id int) (int64, error) {
	m:=new(models.AreaExt)
	m.Id = id
	num, err := db.Db().Engine.Delete(m)
	if err == nil {
		return num, nil
	}
	return num, err
}