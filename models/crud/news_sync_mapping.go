
package crud

import (
	"fmt"
	"github.com/foxiswho/shop-go/models"
	"github.com/foxiswho/shop-go/module/db"
	"github.com/foxiswho/shop-go/util"
)

type NewsSyncMappingCrud struct {

}

func NewNewsSyncMappingCrud() *NewsSyncMappingCrud {
	return new(NewsSyncMappingCrud)
}

//初始化列表
func news_sync_mappingNewMakeDataArr() []models.NewsSyncMapping {
	return make([]models.NewsSyncMapping, 0)
}

//列表查询
func (s *NewsSyncMappingCrud) GetAll(where []*db.QueryCondition, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
	m := models.NewNewsSyncMapping()
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
	data := news_sync_mappingNewMakeDataArr()
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
func (s *NewsSyncMappingCrud) GetById(id int) (*models.NewsSyncMapping, error) {
    m:=new(models.NewsSyncMapping)
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
func (s *NewsSyncMappingCrud) Delete(id int) (int64, error) {
	m:=new(models.NewsSyncMapping)
	m.Id = id
	num, err := db.Db().Engine.Delete(m)
	if err == nil {
		return num, nil
	}
	return num, err
}