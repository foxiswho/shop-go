package admin

import (
	"github.com/foxiswho/shop-go/module/context"
	"net/http"
	"github.com/labstack/echo"
	"github.com/foxiswho/shop-go/module/db"
	"github.com/foxiswho/shop-go/service/admin_service"
	"github.com/foxiswho/shop-go/util/conv"
)

func AdminListHandler(c *context.BaseContext) error {
	idStr := c.Param("id")
	page, _ := conv.ObjToInt(idStr)
	if page < 1 {
		page = 1
	}
	//查询
	where := db.NewMakeQueryCondition();
	where = append(where, db.AddQueryCondition("id", ">", 0))
	data, err := admin_service.GetAll(where, page, 20)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"code":    http.StatusBadRequest,
			"message": "获取数据发生错误",
			"data":    data,
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"code":    http.StatusOK,
		"message": "获取数据",
		"data":   data,
	})
}
