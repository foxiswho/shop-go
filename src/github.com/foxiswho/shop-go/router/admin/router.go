package admin

import (
	"github.com/labstack/echo"
	"github.com/foxiswho/shop-go/module/context"
	. "github.com/foxiswho/shop-go/module/conf"
	"github.com/foxiswho/shop-go/module/session"
	"github.com/foxiswho/shop-go/module/cache"
	mw "github.com/labstack/echo/middleware"
	"github.com/foxiswho/shop-go/middleware/opentracing"
	"net/http"
	"github.com/foxiswho/shop-go/router/admin/login"
	admin2 "github.com/foxiswho/shop-go/router/admin/admin"
	"github.com/foxiswho/shop-go/module/jwt"
)

func RoutersAdmin() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Context自定义
	e.Use(context.NewBaseContext())
	e.Use(context.SetContextTypeAdmin())
	// Customization
	if Conf.ReleaseMode {
		e.Debug = false
	}
	e.Logger.SetPrefix("admin")
	e.Logger.SetLevel(GetLogLvl())

	// Session
	e.Use(session.Session())

	// OpenTracing
	if !Conf.Opentracing.Disable {
		e.Use(opentracing.OpenTracing("admin"))
	}

	// CSRF
	//e.Use(csrf.CSRFWithConfig())

	// Gzip
	e.Use(mw.GzipWithConfig(mw.GzipConfig{
		Level: 5,
	}))

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	//e.Static("/favicon.ico", "./assets/img/favicon.ico")

	// Cache
	e.Use(cache.Cache())

	// Unauthenticated route
	e.GET("/", accessible)
	////////////////////////////
	/////admin
	admin_login := e.Group("/admin_login")
	{
		//admin_login.Use(context.SetContextTypeAdmin())
		//admin_login.Use(admin_auth.New(auth.GenerateAnonymousUser()))
		admin_login.GET("/", accessible)
		admin_login.POST("/login", context.Handler(login.LoginPostHandler))
		admin_login.POST("/logout", context.Handler(login.LogoutPostHandler))
	}
	admin := e.Group("/admin")
	{
		admin.Use(jwt.GetJwtMiddlewareAdmin())

		admin.GET("/admin/info", context.Handler(admin2.AdminInfoGetHandler))
		admin.GET("/admin", context.Handler(admin2.AdminListHandler))
		admin.PUT("/admin", context.Handler(admin2.AdminPutHandler))
		//admin.GET("/index", context.Handler(api.JwtTesterApiHandler))
	}
	return e
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}
