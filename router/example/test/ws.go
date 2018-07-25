package test

import (
	"github.com/foxiswho/shop-go/module/context"
)

func WsHandler(c *context.BaseContext) error {
	c.Set("tmpl", "example/test/ws")
	c.Set("data", map[string]interface{}{
		"title": "Web Socket",
	})
	return nil
}
