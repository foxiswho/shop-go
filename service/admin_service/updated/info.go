package updated

import (
	"github.com/foxiswho/shop-go/models"
	"github.com/foxiswho/shop-go/util"
	"github.com/foxiswho/shop-go/module/db"
)

type UpdateInfo struct {
	admin *models.Admin
}

func NewUpdateInfo() *UpdateInfo {
	return new(UpdateInfo)
}

func (c *UpdateInfo) SetAdmin(admin *models.Admin) {
	c.admin = admin
}

func (c *UpdateInfo) check() (bool, error) {
	if c.admin == nil {
		return false, util.NewError("数据 没有传入")
	}
	if len(c.admin.Username) < 1 {
		return false, util.NewError("用户名 不能为空")
	}
	if len(c.admin.Mail) < 1 {
		return false, util.NewError("邮箱 不能为空")
	}
	if len(c.admin.Name) < 1 {
		return false, util.NewError("显示名称 不能为空")
	}
	return true, nil
}

func (c *UpdateInfo) Process() (bool, error) {
	//只更新指定字段
	fields := []string{"username", "mail", "job_no", "nick_name", "true_name", "qq", "phone", "mobile", "name"}
	ok, err := c.check()
	if err != nil {
		return ok, err
	}
	_, err = db.Db().Engine.Id(c.admin.Id).Cols(fields...).Update(c.admin)
	if err != nil {
		return false, err
	}
	return true, nil
}
