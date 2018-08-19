package index

import "github.com/foxiswho/shop-go/module/context"

func IndexHandler(c *context.BaseContext) error {
	c.Set("tmpl", "example2/index/home")
	c.Set("data", map[string]interface{}{
		"title": "Home",
	})

	return nil
}