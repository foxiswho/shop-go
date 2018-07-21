package admin

import (
	"github.com/foxiswho/shop-go/module/context"

	"github.com/labstack/echo"
	"net/http"
	"fmt"
)

func AdminGetHandler(c *context.BaseContext) error {
	claims := c.JwtTokenGetAdmin()
	if claims == nil {
		return echo.ErrUnauthorized
	}
	fmt.Println("admin.Claims id=", claims["id"])
	//fmt.Println("admin.Claims", maps["id"])
	return c.JSON(http.StatusOK, echo.Map{
		"code":    http.StatusOK,
		"message": "获取数据",
		"name":    "tmp",
		"avatar":  "/",
		"roles":   []int{1},
	})
	return echo.ErrUnauthorized
}
