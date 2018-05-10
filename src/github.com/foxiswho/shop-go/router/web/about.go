package web

import ()

func AboutHandler(c *Context) error {
	c.Set("tmpl", "web/about")
	c.Set("data", map[string]interface{}{
		"title": "About",
	})

	return nil
}
