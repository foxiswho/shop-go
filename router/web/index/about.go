package index

import (
	"github.com/foxiswho/shop-go/module/context"
)

func AboutHandler(c *context.BaseContext) error {
	c.Set("tmpl", "web/index/about")
	c.Set("data", map[string]interface{}{
		"title": "About",
	})

	return nil
}
