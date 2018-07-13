package auth

import (
	"github.com/labstack/echo"
	"github.com/foxiswho/shop-go/router/base"
	"github.com/foxiswho/shop-go/consts/context"
	"github.com/foxiswho/shop-go/module/auth/admin"
	"github.com/foxiswho/shop-go/module/auth/user_auth"
)

func GetAuthData(c echo.Context) interface{} {
	context_type := base.GetContextType(c)
	if context.Type_Admin == context_type {
		return admin.DefaultGetAdmin(c)
	} else {
		return user_auth.DefaultGetUser(c)
	}
}
