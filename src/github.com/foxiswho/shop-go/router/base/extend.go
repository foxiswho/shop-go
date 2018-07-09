package base

import (
	"github.com/labstack/echo"
	"github.com/foxiswho/shop-go/consts/context"
)

//设置 为管理员
func SetContextTypeAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c BaseContext) error {
			c.ContextType = context.Type_Admin
			return next(c)
		}
	}
}

//设置 为用户
func SetContextTypeUser() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c BaseContext) error {
			c.ContextType = context.Type_Admin
			return next(c)
		}
	}
}

func GetContextType(c BaseContext) string {
	return c.ContextType
}
