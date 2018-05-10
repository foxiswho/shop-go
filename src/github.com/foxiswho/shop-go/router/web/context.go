package web

import (
	"github.com/labstack/echo"

	"github.com/foxiswho/shop-go/middleware/session"

	"github.com/opentracing/opentracing-go"

	"github.com/foxiswho/shop-go/module/auth"
	ot "github.com/foxiswho/shop-go/middleware/opentracing"
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

func (ctx *Context) Session() session.Session {
	return session.Default(ctx)
}

func (ctx *Context) Auth() auth.Auth {
	return auth.Default(ctx)
}

func (ctx *Context) OpenTracingSpan() opentracing.Span {
	return ot.Default(ctx)
}
