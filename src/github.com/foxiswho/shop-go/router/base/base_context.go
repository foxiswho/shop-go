package base

import (
	"github.com/labstack/echo"

	"github.com/foxiswho/shop-go/middleware/session"

	"github.com/opentracing/opentracing-go"

	"github.com/foxiswho/shop-go/module/auth"
	ot "github.com/foxiswho/shop-go/middleware/opentracing"
)

type BaseContext struct {
	echo.Context
}

func NewBaseContext() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := &BaseContext{c}
			return h(ctx)
		}
	}
}

func (ctx *BaseContext) Session() session.Session {
	return session.Default(ctx)
}

func (ctx *BaseContext) Auth() auth.Auth {
	return auth.Default(ctx)
}

func (ctx *BaseContext) OpenTracingSpan() opentracing.Span {
	return ot.Default(ctx)
}
