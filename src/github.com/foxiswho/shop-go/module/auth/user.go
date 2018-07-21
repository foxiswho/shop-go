package auth

import (
	"github.com/foxiswho/shop-go/module/auth/user_auth"
	"github.com/foxiswho/shop-go/module/auth/admin_auth"
	"context"
)

//user
func AuthUser(c *context.Context) user_auth.AuthUser {
	return user_auth.Default(c)
}

//admin 后台
func AuthAdmin(c *context.Context) admin_auth.AuthAdmin {
	return admin_auth.Default(c)
}
