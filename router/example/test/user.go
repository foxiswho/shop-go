package test

import (
	"strconv"

	"github.com/foxiswho/shop-go/service/example_service"
	"fmt"
	"github.com/foxiswho/shop-go/module/context"
)

func UserHandler(c *context.BaseContext) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println("idStr=>", idStr)
	fmt.Println("id=>", id)
	u := example_service.GetUserById(id)
	fmt.Println("UserHandler", u)
	c.Set("tmpl", "web/user_service")
	c.Set("data", map[string]interface{}{
		"title":        "Admin",
		"user_service": u,
	})

	return nil
}
