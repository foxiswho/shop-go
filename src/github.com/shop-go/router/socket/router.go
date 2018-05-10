package socket

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/hb-go/echo-web/model"
	"github.com/hb-go/echo-web/module/auth"
	"github.com/hb-go/echo-web/module/cache"
	"github.com/hb-go/echo-web/module/render"
	"github.com/hb-go/echo-web/module/session"
)

func Routers() *echo.Echo {
	e := echo.New()

	// Session
	e.Use(session.Session())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 模板
	e.Renderer = render.LoadTemplates()
	e.Use(render.Render())

	// Cache
	e.Use(cache.Cache())

	// Auth
	e.Use(auth.New(model.GenerateAnonymousUser))

	e.GET("/ws", socketHandler)

	return e
}
