package api

import (
	"github.com/hb-go/json"
	"github.com/labstack/echo"
	"github.com/opentracing/opentracing-go"

	ot "github.com/hb-go/echo-web/middleware/opentracing"
)

func NewContext() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := &Context{c}
			return h(ctx)
		}
	}
}

type Context struct {
	echo.Context
}

/**
 * 以接口参数或后缀返回数据
 * JSONP、JSON/XML
 */
func (c *Context) AutoFMT(code int, i interface{}) (err error) {
	// JSONP
	callback := c.QueryParam("jsonp")
	if len(callback) > 0 {
		c.Logger().Infof("JSONP callback func:%v", callback)
		return c.JSONP(code, callback, i)
	} else {
		return c.JSON(code, i)
	}
}

// 自定义JSON解析
func (c *Context) CustomJSON(code int, i interface{}, f string) (err error) {
	if c.Context.Echo().Debug {
		return c.JSONPretty(code, i, "  ")
	}
	b, err := json.MarshalFilter(i, f)
	if err != nil {
		return
	}
	return c.JSONBlob(code, b)
}

func (ctx *Context) OpenTracingSpan() opentracing.Span {
	return ot.Default(ctx)
}
