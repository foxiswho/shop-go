package design

import (
	"net/http"
	"github.com/foxiswho/shop-go/module/context"
	"github.com/labstack/echo"
	"github.com/foxiswho/shop-go/module/log"
	"github.com/foxiswho/shop-go/module/design/service"
)

func ServiceMakeHandler(c *context.BaseContext) error {
	template_file := "./template/design/make/service.go.tpl"
	json := make(map[string]interface{})
	c.FormJson(json)
	log.Debugf(" input dir: %v", json)
	service_path := json["dir"].(string)
	log.Debugf(" input dir: %v", service_path)
	if len(service_path) < 1 {
		service_path = "./models/crud"
	}
	err := service.MakeService(template_file, service_path)
	if err != nil {
		return c.JSON(http.StatusOK, echo.Map{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
	} else {
		return c.JSON(http.StatusOK, echo.Map{
			"code":    http.StatusOK,
			"message": "生成成功",
		})
	}
}