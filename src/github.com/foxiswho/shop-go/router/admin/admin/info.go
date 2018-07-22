package admin

import (
	"github.com/foxiswho/shop-go/module/context"

	"github.com/labstack/echo"
	"net/http"
	"fmt"
)

func AdminInfoGetHandler(c *context.BaseContext) error {
	fmt.Println("admin.Claims id=", c.GetUserId())
	//fmt.Println("admin.Claims", maps["id"])
	return c.JSON(http.StatusOK, echo.Map{
		"code":         http.StatusOK,
		"message":      "获取数据",
		"name":         "管理员",
		"introduction": "我是超级管理员",
		"avatar":       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		"roles":        []string{"admin"},
	})
	return echo.ErrUnauthorized
}
