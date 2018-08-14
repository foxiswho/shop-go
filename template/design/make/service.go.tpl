
package crud

import (
	"fmt"
	"github.com/foxiswho/shop-go/models"
	"github.com/foxiswho/shop-go/module/db"
	"github.com/foxiswho/shop-go/util"
)

type {{.tables_Camel_Case}}Crud struct {

}

func New{{.tables_Camel_Case}}Crud() *{{.tables_Camel_Case}}Crud {
	return new({{.tables_Camel_Case}}Crud)
}

//初始化列表
func {{.tables}}NewMakeDataArr() []models.{{.tables_Camel_Case}} {
	return make([]models.{{.tables_Camel_Case}}, 0)
}

//列表查询
func (s *{{.tables_Camel_Case}}Crud) GetAll(where []*db.QueryCondition, fields []string, orderBy string, page int, limit int) (*db.Paginator, error) {
	m := models.New{{.tables_Camel_Case}}()
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
	data := {{.tables}}NewMakeDataArr()
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
func (s *{{.tables_Camel_Case}}Crud) GetById(id int) (*models.{{.tables_Camel_Case}}, error) {
    m:=new(models.{{.tables_Camel_Case}})
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
func (s *{{.tables_Camel_Case}}Crud) GetByIds(id []int) ([]*models.{{.tables_Camel_Case}}, error) {
    data := {{.tables}}NewMakeDataArr()
	err := db.Db().Engine.In("id", id).Find(&data)
    if err != nil {
        return nil, err
    }
    return data, nil
}

// 获取 多条记录
func (s *{{.tables_Camel_Case}}Crud) GetByIdsIndex(id []int) (map[int]*models.{{.tables_Camel_Case}}, error) {
    data := {{.tables}}NewMakeDataArr()
	err := db.Db().Engine.In("id", id).Find(&data)
    if err != nil {
        return nil, err
    }
    data_index := make(map[int]*models.{{.tables_Camel_Case}})
    for _, val := range data {
        data_index[val.Id] = &val
    }
    return data_index, nil
}

// 删除 单条记录
func (s *{{.tables_Camel_Case}}Crud) Delete(id int) (int64, error) {
	m:=new(models.{{.tables_Camel_Case}})
	m.Id = id
	num, err := db.Db().Engine.Delete(m)
	if err == nil {
		return num, nil
	}
	return num, err
}