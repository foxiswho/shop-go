package context

import (
	"github.com/foxiswho/shop-go/module/auth/user_auth"
	"github.com/foxiswho/shop-go/module/auth/admin_auth"
)

//user
func (c *BaseContext) AuthUser() user_auth.AuthUser {
	return user_auth.Default(c)
}

//admin 后台
func (c *BaseContext) AuthAdmin() admin_auth.AuthAdmin {
	return admin_auth.Default(c)
}
