package context

import (
	"github.com/labstack/echo"
	"net/http"
)

type (
	HandlerFunc func(*BaseContext) error
)

/**
 * 自定义Context的Handler
 */
func Handler(h HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(*BaseContext)
		return h(ctx)
	}
}

func Accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}
