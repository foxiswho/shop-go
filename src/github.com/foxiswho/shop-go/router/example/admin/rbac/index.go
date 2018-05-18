package rbac

import (
	"github.com/foxiswho/shop-go/router/base"
)

func IndexHandler(c *base.BaseContext) error {
	c.Set("tmpl", "example/admin/rbac/index")
	c.Set("data", map[string]interface{}{
		"title": "rbac",
	})

	return nil
}
