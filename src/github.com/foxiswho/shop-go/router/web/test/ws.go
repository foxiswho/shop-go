package test

import (
	"github.com/foxiswho/shop-go/router/web"
)

func WsHandler(c *web.BaseContext) error {
	c.Set("tmpl", "web/ws")
	c.Set("data", map[string]interface{}{
		"title": "Web Socket",
	})
	return nil
}
