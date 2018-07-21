package context

import (
	"github.com/labstack/echo"
	"github.com/foxiswho/shop-go/consts/context"
)

//设置 为管理员
func SetContextTypeAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			x := c.(*BaseContext)
			x.ContextType = context.Type_Admin
			return next(x)
		}
	}
}

//设置 为用户
func SetContextTypeUser() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			x := c.(*BaseContext)
			x.ContextType = context.Type_Admin
			return next(x)
		}
	}
}

// 获取 用户登录类别 user  admin
func GetContextType(c echo.Context) string {
	x := c.(*BaseContext)
	return x.ContextType
}

//设置会话 为 jwt方式
func SetSessionTypeJwt() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			x := c.(*BaseContext)
			x.SessionType = context.Session_jwt
			return next(x)
		}
	}
}

//设置会话 为  cookie
func SetSessionTypeCookie() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			x := c.(*BaseContext)
			x.SessionType = context.Session_cookie
			return next(x)
		}
	}
}

// 获取用户号会话类别
func GetSessionType(c echo.Context) string {
	x := c.(*BaseContext)
	return x.SessionType
}
