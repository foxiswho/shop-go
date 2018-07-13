package test

import (
	"net/http"
	"github.com/foxiswho/shop-go/module/auth/user_auth"
	"github.com/foxiswho/shop-go/module/context"
)

func LogoutHandler(c *context.BaseContext) error {
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
