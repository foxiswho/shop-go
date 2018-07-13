package test

import (
	"net/http"

	"github.com/foxiswho/shop-go/module/auth/user_auth"
	"github.com/foxiswho/shop-go/module/log"
	userService "github.com/foxiswho/shop-go/service/example_service"
	"github.com/foxiswho/shop-go/router/base"
)

func RegisterHandler(c *base.BaseContext) error {
	redirect := c.QueryParam(user_auth.RedirectParam)

	a := c.Auth()
	if a.User.IsAuthenticated() {
		if redirect == "" {
			redirect = "/"
		}
		c.Redirect(http.StatusMovedPermanently, redirect)
		return nil
	}

	c.Set("tmpl", "example/test/register")
	c.Set("data", map[string]interface{}{
		"title":         "Register",
		"redirectParam": user_auth.RedirectParam,
		"redirect":      redirect,
	})

	return nil
}

func RegisterPostHandler(c *base.BaseContext) error {
	redirect := c.QueryParam(user_auth.RedirectParam)
	if redirect == "" {
		redirect = "/"
	}

	a := c.Auth()
	if a.User.IsAuthenticated() {
		c.Redirect(http.StatusMovedPermanently, redirect)
		return nil
	}

	var form LoginForm
	if err := c.Bind(&form); err == nil {
		u := userService.AddUserWithNicknamePwd(form.Nickname, form.Password)
		if u != nil {
			session := c.Session()
			err := user_auth.AuthenticateSession(session, u)
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
			}
			c.Redirect(http.StatusMovedPermanently, redirect)
			return nil
		} else {
			log.Debugf("Register user_service add error")

			s := c.Session()
			s.AddFlash("Register user_service add error", "_error")

			// registerURL := c.Request().URI()
			// c.Redirect(http.StatusMovedPermanently, registerURL)
			c.Set("tmpl", "web/register")
			c.Set("data", map[string]interface{}{
				"title":         "Register",
				"redirectParam": user_auth.RedirectParam,
				"redirect":      redirect,
			})
			return nil
		}
	} else {
		log.Debugf("Register form bind Error: %v", err)

		s := c.Session()
		s.AddFlash("Register form bind Error:"+err.Error(), "_error")

		// registerURL := c.Request().URI()
		// c.Redirect(http.StatusMovedPermanently, registerURL)
		c.Set("tmpl", "example/test/register")
		c.Set("data", map[string]interface{}{
			"title":         "Register",
			"redirectParam": user_auth.RedirectParam,
			"redirect":      redirect,
		})
		return nil
	}
}
