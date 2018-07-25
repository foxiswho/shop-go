package admin

import (
	"github.com/foxiswho/shop-go/module/context"
	"net/http"
	"github.com/labstack/echo"
	"github.com/foxiswho/shop-go/models"
	"github.com/foxiswho/shop-go/service/admin_service/updated"
)

//更新数据
func AdminPutHandler(c *context.BaseContext) error {
	admin := models.NewAdmin()
	if err := c.Bind(admin); err == nil {
		a := updated.NewUpdateInfo()
		a.SetAdmin(admin)
		ok, err := a.Process()
		if ok {
			return c.JSON(http.StatusOK, echo.Map{
				"code":    http.StatusOK,
				"message": "操作成功",
			})

		} else {
			return c.JSON(http.StatusOK, echo.Map{
				"code":    http.StatusBadRequest,
				"message": err.Error(),
			})
		}
	}
	return c.JSON(http.StatusUnauthorized, echo.Map{
		"code":    http.StatusUnauthorized,
		"message": "数据错误",
	})
}
