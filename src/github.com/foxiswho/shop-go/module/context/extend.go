package context

import (
	"github.com/labstack/echo"
	"github.com/foxiswho/shop-go/module/auth/user_auth"
	"github.com/foxiswho/shop-go/module/auth/admin_auth"
)

//设置 为管理员
func SetContextTypeAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			x := c.(*BaseContext)
			x.ContextType = Type_Admin
			return next(x)
		}
	}
}

//设置 为用户
func SetContextTypeUser() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			x := c.(*BaseContext)
			x.ContextType = Type_Admin
			return next(x)
		}
	}
}

func GetContextType(c echo.Context) string {
	x := c.(*BaseContext)
	return x.ContextType
}

//user
func GetAuthUser(c echo.Context) user_auth.AuthUser {
	return user_auth.Default(c)
}

//admin 后台
func GetAuthAdmin(c echo.Context) admin_auth.AuthAdmin {
	return admin_auth.Default(c)
}
