package design

import (
	"github.com/foxiswho/shop-go/module/context"
	"net/http"
	"github.com/labstack/echo"
	"fmt"
	"github.com/foxiswho/shop-go/module/design/models"
	"github.com/foxiswho/shop-go/module/log"
)

func ModelsMakeHandler(c *context.BaseContext) error {

	template_file := "./template/design/make/models.go.tpl"
	fmt.Println(template_file)
	json:=make(map[string]interface{})
	c.FormJson(json)
	log.Debugf(" input dir: %v", json)
	path := json["dir"].(string)
	log.Debugf(" input dir: %v", path)
	if len(path) < 1 {
		path = "models"
	}
	models.MakeModels(template_file,path)


	return c.JSON(http.StatusOK, echo.Map{
		"code":    http.StatusOK,
		"message": "生成成功",
	})
}