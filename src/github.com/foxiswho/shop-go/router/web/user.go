package web

import (
	"strconv"

	"github.com/foxiswho/shop-go/service/user"
	"fmt"
)

func UserHandler(c *Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println("idStr=>",idStr)
	fmt.Println("id=>",id)
	u := user.GetUserById(id)
	fmt.Println("UserHandler",u)
	c.Set("tmpl", "web/user")
	c.Set("data", map[string]interface{}{
		"title": "User",
		"user":  u,
	})

	return nil
}
