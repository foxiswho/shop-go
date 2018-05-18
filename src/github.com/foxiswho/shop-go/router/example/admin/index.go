package admin

import (
	"github.com/foxiswho/shop-go/router/base"
)

func IndexHandler(c *base.BaseContext) error {
	c.Set("tmpl", "example/admin/index")
	c.Set("data", map[string]interface{}{
		"title": "Home",
	})

	return nil
}
