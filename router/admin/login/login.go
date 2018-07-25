package login

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/foxiswho/shop-go/module/context"
	"github.com/foxiswho/shop-go/service/admin_service"
	"github.com/foxiswho/shop-go/module/log"
)

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// 登录
func LoginPostHandler(c *context.BaseContext) error {
	var form LoginForm
	if err := c.Bind(&form); err == nil {
		log.Debugf("post form :%v", form)
		token, err := admin_service.Login(form.Username, form.Password)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"code": http.StatusUnauthorized,
				"message": err.Error(),
			})

		} else {
			return c.JSON(http.StatusOK, echo.Map{
				"code":   http.StatusOK,
				"message":   "登录成功",
				"token": token,
			})
		}
	} else {
		params, _ := c.FormParams()
		log.Debugf("Login form params: %v", params)
		log.Debugf("Login form bind Error: %v", err)
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"code":     http.StatusUnauthorized,
			"message": "错误",
		})
	}
	return echo.ErrUnauthorized
}
