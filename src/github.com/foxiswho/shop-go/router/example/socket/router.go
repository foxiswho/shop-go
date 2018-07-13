package socket

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/foxiswho/shop-go/module/auth/user_auth"
	"github.com/foxiswho/shop-go/module/cache"
	"github.com/foxiswho/shop-go/module/render"
	"github.com/foxiswho/shop-go/module/session"
	authService "github.com/foxiswho/shop-go/service/user_service/auth"
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

	// AuthUser
	e.Use(user_auth.New(authService.GenerateAnonymousUser))

	e.GET("/ws", socketHandler)

	return e
}
