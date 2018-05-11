package user

import (
	"net/http"

	"github.com/dchest/captcha"

	"github.com/foxiswho/shop-go/module/auth"
	"github.com/foxiswho/shop-go/module/log"
	"fmt"
	userService "github.com/foxiswho/shop-go/service/user_service"
	"github.com/foxiswho/shop-go/router/base"
)

type LoginForm struct {
	Nickname string `form:"nickname" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func LoginHandler(c *base.BaseContext) error {
	redirect := c.QueryParam(auth.RedirectParam)

	a := c.Auth()
	if a.User.IsAuthenticated() {
		if redirect == "" {
			redirect = "/"
		}
		c.Redirect(http.StatusMovedPermanently, redirect)
		return nil
	}

	c.Set("tmpl", "web/login")
	c.Set("data", map[string]interface{}{
		"title":         "Login",
		"redirectParam": auth.RedirectParam,
		"redirect":      redirect,
		"CaptchaId":     captcha.NewLen(6),
	})

	return nil
}

func LoginPostHandler(c *base.BaseContext) error {
	loginURL := c.Request().RequestURI

	if !captcha.VerifyString(c.FormValue("captchaId"), c.FormValue("captchaSolution")) {
		log.Debugf("Wrong captcha solution: %v! No robots allowed!\n", c.Param("captchaSolution"))
		c.Redirect(http.StatusMovedPermanently, loginURL)
		return nil
	} else {
		log.Debugf("Great job, human! You solved the captcha.")
	}

	redirect := c.QueryParam(auth.RedirectParam)
	if redirect == "" {
		redirect = "/"
	}

	a := c.Auth()
	if a.User.IsAuthenticated() {
		fmt.Println("已经验证过了")
		c.Redirect(http.StatusMovedPermanently, redirect)
		return nil
	}

	var form LoginForm
	if err := c.Bind(&form); err == nil {
		fmt.Println("form",form)
		u := userService.GetUserByNicknamePwd(form.Nickname, form.Password)
		fmt.Println("db=>u")
		fmt.Println("db=>u")
		fmt.Println("db=>u")
		fmt.Println("db=>u")
		fmt.Println("db=>u")
		fmt.Println("db=>u",u)
		if u != nil {
			session := c.Session()
			err := auth.AuthenticateSession(session, u)
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
			}
			c.Redirect(http.StatusMovedPermanently, redirect)
			return nil
		} else {
			c.Redirect(http.StatusMovedPermanently, loginURL)
			return nil
		}
	} else {
		params, _ := c.FormParams()
		log.Debugf("Login form params: %v", params)
		log.Debugf("Login form bind Error: %v", err)
		c.Redirect(http.StatusMovedPermanently, loginURL)
		return nil
	}
}

func LogoutHandler(c *base.BaseContext) error {
	session := c.Session()
	a := c.Auth()
	auth.Logout(session, a.User)

	redirect := c.QueryParam(auth.RedirectParam)
	if redirect == "" {
		redirect = "/"
	}

	c.Redirect(http.StatusMovedPermanently, redirect)

	return nil
}

func RegisterHandler(c *base.BaseContext) error {
	redirect := c.QueryParam(auth.RedirectParam)

	a := c.Auth()
	if a.User.IsAuthenticated() {
		if redirect == "" {
			redirect = "/"
		}
		c.Redirect(http.StatusMovedPermanently, redirect)
		return nil
	}

	c.Set("tmpl", "web/register")
	c.Set("data", map[string]interface{}{
		"title":         "Register",
		"redirectParam": auth.RedirectParam,
		"redirect":      redirect,
	})

	return nil
}

func RegisterPostHandler(c *base.BaseContext) error {
	redirect := c.QueryParam(auth.RedirectParam)
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
			err := auth.AuthenticateSession(session, u)
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
				"redirectParam": auth.RedirectParam,
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
		c.Set("tmpl", "web/register")
		c.Set("data", map[string]interface{}{
			"title":         "Register",
			"redirectParam": auth.RedirectParam,
			"redirect":      redirect,
		})
		return nil
	}
}
