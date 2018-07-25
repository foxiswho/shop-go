package test

import (
	"github.com/foxiswho/shop-go/module/context"
)

func JwtLoginHandler(c *context.BaseContext) error {
	c.Set("tmpl", "example/test/jwt_login")
	c.Set("data", map[string]interface{}{
		"title": "JWT 接口测试",
	})

	return nil
}
