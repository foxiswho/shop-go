package admin_service

import (
	"github.com/foxiswho/shop-go/service"
	"github.com/foxiswho/shop-go/module/db"
	"github.com/foxiswho/shop-go/models"
)

//列表数据
func GetAll(where []*db.QueryCondition, page, limit int) (*db.Paginator, error) {
	fields := []string{}
	orderBy := ""
	s := service.NewAdminService()
	Query, err := s.GetAll(where, fields, orderBy, page, limit)
	if err != nil {
		return nil, err
	}
	if Query.TotalCount > 0 {
		for y, x := range Query.Data {
			//密码 salt 清空 不显示
			admin := x.(models.Admin)
			admin.Password = ""
			admin.Salt = ""
			Query.Data[y] = admin
		}
	}
	return Query, nil
}
