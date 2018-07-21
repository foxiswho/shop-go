package auth_middleware

import (
	"github.com/labstack/echo"
	"github.com/foxiswho/shop-go/module/context"
	"github.com/foxiswho/shop-go/module/auth/admin_auth"
	"github.com/foxiswho/shop-go/module/auth/user_auth"
	"github.com/foxiswho/shop-go/models/auth"
)

func NewAdmin(newAdmin func() auth.User) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(x echo.Context) error {
			c := x.(*context.BaseContext)
			//userId := session_type.GetUserId(c)
			userId := c.GetUserId()
			user := newAdmin()
			if userId > 0 {
				err := user.GetById(userId)
				if err != nil {
					c.Logger().Errorf("Login Error: %v", err)
				} else {
					user.Login()
				}
			} else {
				c.Logger().Debugf("Login status: No UserId")
			}
			auth := admin_auth.AuthAdmin{user}
			c.Set(admin_auth.DefaultKey, auth)
			return next(c)
		}
	}
}

func NewUser(newUser func() auth.User) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(x echo.Context) error {
			c := x.(*context.BaseContext)
			//userId := session_type.GetUserId(c)
			userId := c.GetUserId()
			user := newUser()
			if userId > 0 {
				err := user.GetById(userId)
				if err != nil {
					c.Logger().Errorf("Login Error: %v", err)
				} else {
					user.Login()
				}
			} else {
				c.Logger().Debugf("Login status: No UserId")
			}
			auth := user_auth.AuthUser{user}
			c.Set(user_auth.DefaultKey, auth)
			return next(c)
		}
	}
}
