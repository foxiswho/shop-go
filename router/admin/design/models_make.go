package design

import (
	"github.com/foxiswho/shop-go/module/context"
	"net/http"
	"github.com/labstack/echo"
	"os/exec"
	"github.com/foxiswho/shop-go/module/log"
	"github.com/foxiswho/shop-go/module/conf"
	"fmt"
)

func ModelsMakeHandler(c *context.BaseContext) error {
	return c.JSON(http.StatusOK, echo.Map{
		"code":    http.StatusBadRequest,
		"message": "此功能 未完成",
	})
	//
	json:=make(map[string]interface{})
	c.FormJson(json)
	log.Debugf(" FORM JSON: %v", json)

	command := " reverse mysql "+conf.Conf.DB.UserName+":"+conf.Conf.DB.Pwd+"@/"+conf.Conf.DB.Name+"?charset=utf8 template/design/goxorm"
	//
	log.Debugf("Command: xorm %v",command)
	fmt.Println("xorm",command)
	cmd := exec.Command("xorm", command)
	bytes,err := cmd.Output()
	if err != nil {
		log.Debugf("Command Error: %v",err)
		return c.JSON(http.StatusOK, echo.Map{
			"code":    http.StatusBadRequest,
			"message": " 生成失败 error:" + err.Error(),
		})
	}
	resp := string(bytes)
	log.Debugf("Command RESULT: %v",resp)
	return c.JSON(http.StatusOK, echo.Map{
		"code":    http.StatusOK,
		"message": "生成成功",
	})
}