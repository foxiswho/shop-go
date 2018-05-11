package index

import (
	"github.com/foxiswho/shop-go/router/web"
)

func AboutHandler(c *web.BaseContext) error {
	c.Set("tmpl", "web/about")
	c.Set("data", map[string]interface{}{
		"title": "About",
	})

	return nil
}
