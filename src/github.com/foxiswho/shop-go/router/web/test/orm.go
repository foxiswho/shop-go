package test

import (
	"github.com/foxiswho/shop-go/router/base"
	"github.com/foxiswho/shop-go/service/user_service"
	"fmt"
	"github.com/foxiswho/shop-go/util/str"
	"github.com/foxiswho/shop-go/module/db"
)

type Orm struct {
}

func NewOrm() *Orm {
	return new(Orm)
}

func (x *Orm) IndexHandler(c *base.BaseContext) error {
	//主键ID查询
	user := user_service.GetUserById(1)
	fmt.Println("GetUserById user=>", user)
	//sql语句查询并验证
	u := user_service.GetUserByNicknamePwd("admin", "111111")
	fmt.Println("GetUserByNicknamePwd user=>", u)
	//添加用户
	name := "admin_" + str.RandSalt()
	fmt.Println("name=>", name)
	add := user_service.AddUserWithNicknamePwd(name, "111111")
	fmt.Println("AddUserWithNicknamePwd user=>", add)
	//查询
	where := db.NewMakeQueryCondition();
	where = append(where, db.AddQueryCondition("id", ">", 0))
	data, err := user_service.GetAll(where, []string{}, "id desc", 1, 20)
	fmt.Println("GetAll err=>", err)
	fmt.Println("GetAll data=>", data)
	c.Set("tmpl", "web/test/orm")
	c.Set("data", map[string]interface{}{
		"title": "测试 COOIE",
	})
	return nil
}
