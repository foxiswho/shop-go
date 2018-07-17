package auth

import (
	"github.com/labstack/echo"
	"github.com/foxiswho/shop-go/module/auth/admin_auth"
	"github.com/foxiswho/shop-go/module/auth/user_auth"
	"github.com/foxiswho/shop-go/module/context"
)

func GetAuthData(c echo.Context) interface{} {
	context_type := context.GetContextType(c)
	if context.Type_Admin == context_type {
		return admin_auth.DefaultGetAdmin(c)
	} else {
		return user_auth.DefaultGetUser(c)
	}
}

func GetAuthDataRoleId(c echo.Context) int {
	context_type := context.GetContextType(c)
	if context.Type_Admin == context_type {
		return admin_auth.GetRoleId(c)
	} else {
		return user_auth.GetRoleId(c)
	}
}

func GetAuthDataRoleExtend(c echo.Context) []int {
	context_type := context.GetContextType(c)
	if context.Type_Admin == context_type {
		return admin_auth.GetRoleExtend(c)
	} else {
		return user_auth.GetRoleExtend(c)
	}
}
