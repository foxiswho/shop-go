package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/foxiswho/shop-go/module/cache"
	"github.com/foxiswho/shop-go/module/render"
	"github.com/foxiswho/shop-go/module/session"
	"github.com/casbin/casbin"
	casbinmw "github.com/labstack/echo-contrib/casbin"
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
	ce := casbin.NewEnforcer("casbin_auth_model.conf", "")
	ce.AddRoleForUser("alice", "admin")
	e.Use(casbinmw.MiddlewareWithConfig(casbinmw.Config{Enforcer:ce}))

	return e
}
