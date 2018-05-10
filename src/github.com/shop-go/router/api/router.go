package api

import (
	// "time"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"

	. "github.com/hb-go/echo-web/conf"
	"github.com/hb-go/echo-web/middleware/opentracing"
	"github.com/hb-go/echo-web/module/cache"
	"github.com/hb-go/echo-web/module/session"
)

//-----
// API Routers
//-----
func Routers() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Context自定义
	e.Use(NewContext())

	// Customization
	if Conf.ReleaseMode {
		e.Debug = false
	}
	e.Logger.SetPrefix("api")
	e.Logger.SetLevel(GetLogLvl())

	// Session
	e.Use(session.Session())

	// OpenTracing
	if !Conf.Opentracing.Disable {
		e.Use(opentracing.OpenTracing("api"))
	}

	// CSRF
	e.Use(mw.CSRFWithConfig(mw.CSRFConfig{
		TokenLookup: "form:X-XSRF-TOKEN",
	}))

	// Gzip
	e.Use(mw.GzipWithConfig(mw.GzipConfig{
		Level: 5,
	}))

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	e.Static("/favicon.ico", "./assets/img/favicon.ico")



	// Cache
	e.Use(cache.Cache())

	// e.Use(ec.SiteCache(ec.NewMemcachedStore([]string{conf.MEMCACHED_SERVER}, time.Hour), time.Minute))
	// e.GET("/user/:id", ec.CachePage(ec.NewMemcachedStore([]string{conf.MEMCACHED_SERVER}, time.Hour), time.Minute, UserHandler))

	// Routers
	e.GET("/login", UserLoginHandler)
	e.GET("/register", UserRegisterHandler)

	// 示范github.com/hb-go/json
	e.GET("/json/encode", handler(JsonEncodeHandler))

	// JWT
	r := e.Group("")
	r.Use(mw.JWTWithConfig(mw.JWTConfig{
		SigningKey:  []byte("secret"),
		ContextKey:  "_user",
		TokenLookup: "header:" + echo.HeaderAuthorization,
	}))

	r.GET("/", handler(ApiHandler))

	// curl http://echo.api.localhost:8080/restricted/user -H "Authorization: Bearer XXX"
	r.GET("/user", UserHandler)

	post := r.Group("/post")
	{
		post.GET("/save", PostSaveHandler)
		post.GET("/id/:id", PostHandler)
		post.GET("/:userId/p/:p/s/:s", PostsHandler)
	}

	return e
}

type (
	HandlerFunc func(*Context) error
)

/**
 * 自定义Context的Handler
 */
func handler(h HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(*Context)
		return h(ctx)
	}
}
