package test

import (
	"github.com/foxiswho/shop-go/module/context"
	"fmt"
	"net/http"
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
	"github.com/foxiswho/shop-go/module/log"
	"github.com/foxiswho/shop-go/service/example_service"
	jwt2 "github.com/foxiswho/shop-go/module/jwt"
	jwt3 "github.com/foxiswho/shop-go/consts/session/jwt"
)

func JwtLoginHandler(c *context.BaseContext) error {
	c.Set("tmpl", "example/test/jwt_login")
	c.Set("data", map[string]interface{}{
		"title": "JWT 接口测试",
	})

	return nil
}
// jwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func JwtLoginPostHandler(c *context.BaseContext) error {
	c.Response().Header().Del("Access-Control-Allow-Origin")
	c.Response().Header().Add("Access-Control-Allow-Origin","*")
	var form LoginForm
	if err := c.Bind(&form); err == nil {
		fmt.Println("form",form)
		u := example_service.GetUserByNicknamePwd(form.Nickname, form.Password)
		if u != nil {
			// Generate encoded token and send it as response.
			t, err := jwt2.GetJwtToken(u.Id,jwt3.Jwt_Type_user)
			if err != nil {
				return err
			}
			return c.JSON(http.StatusOK, echo.Map{
				"token": t,
			})
		} else {
			return c.JSON(http.StatusOK, echo.Map{
				"message": "用户不存在",
			})
		}
	} else {
		params, _ := c.FormParams()
		log.Debugf("Login form params: %v", params)
		log.Debugf("Login form bind Error: %v", err)
		return c.JSON(http.StatusOK, echo.Map{
			"message": "错误",
		})
	}
	return echo.ErrUnauthorized
}
