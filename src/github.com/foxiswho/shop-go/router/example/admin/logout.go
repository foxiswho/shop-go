package admin

import (
	"net/http"
	"github.com/foxiswho/shop-go/router/base"
	"github.com/foxiswho/shop-go/module/auth/admin_auth"
)

func LogoutHandler(c *base.BaseContext) error {
	session := c.Session()
	a := c.AuthAdmin()
	admin_auth.Logout(session, a.User)

	redirect := c.QueryParam(admin_auth.RedirectParam)
	if redirect == "" {
		redirect = "/admin_login/login"
	}
	redirect = "/admin_login/login"

	c.Redirect(http.StatusMovedPermanently, redirect)

	return nil
}
