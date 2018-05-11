package web

import ()

func WsHandler(c *BaseContext) error {
	c.Set("tmpl", "web/ws")
	c.Set("data", map[string]interface{}{
		"title": "Web Socket",
	})
	return nil
}
