package admin

import (
	"github.com/foxiswho/shop-go/module/context"
	"fmt"
	"net/http"
	"github.com/labstack/echo"
)
//更新数据
func AdminPutHandler(c *context.BaseContext) error {
	fmt.Println("admin.Claims id=", c.GetUserId())
	//fmt.Println("admin.Claims", maps["id"])
	return c.JSON(http.StatusOK, echo.Map{
		"code":         http.StatusOK,
		"message":      "获取数据",
	})
}
