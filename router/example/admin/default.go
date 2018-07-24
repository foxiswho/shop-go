package admin

import (
	"net/http"
	"github.com/foxiswho/shop-go/module/context"
)

func DefaultHandler(c *context.BaseContext) error {
	c.Redirect(http.StatusMovedPermanently, "/login")
	return nil
}
