package login

import (
	"github.com/foxiswho/shop-go/module/context"
	"net/http"
	"github.com/labstack/echo"
)

//退出
func LogoutPostHandler(c *context.BaseContext) error {
	
	return c.JSON(http.StatusOK, echo.Map{
		"message": "退出成功",
		"code": http.StatusOK,
	})
}
