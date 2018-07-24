package admin

import (
	"net/http"
	"github.com/foxiswho/shop-go/module/auth/admin_auth"
	"github.com/foxiswho/shop-go/module/context"
)

func LogoutHandler(c *context.BaseContext) error {
	session := c.Session()
	a := admin_auth.Default(c)
	admin_auth.Logout(session, a.User)

	redirect := c.QueryParam(admin_auth.RedirectParam)
	if redirect == "" {
		redirect = "/admin_login/login"
	}
	redirect = "/admin_login/login"

	c.Redirect(http.StatusMovedPermanently, redirect)

	return nil
}
