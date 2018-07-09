package render

import (
	"github.com/labstack/echo"
	"github.com/foxiswho/shop-go/router/base"
	"github.com/foxiswho/shop-go/consts/context"
	"github.com/foxiswho/shop-go/module/auth"
	"github.com/foxiswho/shop-go/module/auth_admin"
)

func getAuthData(c echo.Context) interface{} {
	context_type := base.GetContextType(c)
	if context.Type_Admin == context_type {
		return auth_admin.DefaultGetAdmin(c)
	} else {
		return auth.DefaultGetUser(c)
	}
}
