package test

import (
	"github.com/foxiswho/shop-go/router/base"
)

func WsHandler(c *base.BaseContext) error {
	c.Set("tmpl", "web/ws")
	c.Set("data", map[string]interface{}{
		"title": "Web Socket",
	})
	return nil
}
