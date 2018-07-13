package test

import (
	"net/http"
	"github.com/foxiswho/shop-go/router/base"
	"github.com/foxiswho/shop-go/module/auth/user_auth"
)

func LogoutHandler(c *base.BaseContext) error {
	session := c.Session()
	a := c.AuthUser()
	user_auth.Logout(session, a.User)

	redirect := c.QueryParam(user_auth.RedirectParam)
	if redirect == "" {
		redirect = "/"
	}

	c.Redirect(http.StatusMovedPermanently, redirect)

	return nil
}
