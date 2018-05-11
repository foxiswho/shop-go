package user

import (
	"strconv"

	"github.com/foxiswho/shop-go/service/user_service"
	"github.com/foxiswho/shop-go/router/web"
	"fmt"
)

func UserHandler(c *web.BaseContext) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println("idStr=>",idStr)
	fmt.Println("id=>",id)
	u := user_service.GetUserById(id)
	fmt.Println("UserHandler",u)
	c.Set("tmpl", "web/user_service")
	c.Set("data", map[string]interface{}{
		"title": "User",
		"user_service":  u,
	})

	return nil
}
