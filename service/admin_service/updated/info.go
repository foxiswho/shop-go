package updated

import (
	"github.com/foxiswho/shop-go/models"
	"github.com/foxiswho/shop-go/util"
	"github.com/foxiswho/shop-go/module/db"
)

type UpdateInfo struct {
	admin models.Admin
}

func NewUpdateInfo() *UpdateInfo {
	return new(UpdateInfo)
}

func (c *UpdateInfo) SetAdmin(admin models.Admin) {
	c.admin = admin
}

func (c *UpdateInfo) check() (bool, error) {
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
	ok, err := c.check()
	if err != nil {
		return ok, err
	}
	
}
