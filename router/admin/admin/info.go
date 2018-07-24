package admin

import (
	"github.com/foxiswho/shop-go/module/context"

	"github.com/labstack/echo"
	"net/http"
	"fmt"
)

//详情信息
func AdminInfoGetHandler(c *context.BaseContext) error {
	fmt.Println("admin.Claims id=", c.GetUserId())
	//fmt.Println("admin.Claims", maps["id"])
	return c.JSON(http.StatusOK, echo.Map{
		"code":         http.StatusOK,
		"message":      "获取数据",
		"name":         "管理员",
		"introduction": "我是超级管理员",
		"avatar":       "/uploads/image/9358d109b3de9c82bb32fd2d6081800a19d84338.jpg",
		"roles":        []string{"admin"},
	})
	return echo.ErrUnauthorized
}
