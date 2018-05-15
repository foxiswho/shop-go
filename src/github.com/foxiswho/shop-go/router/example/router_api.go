package example

import (

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"

	. "github.com/foxiswho/shop-go/conf"
	"github.com/foxiswho/shop-go/middleware/opentracing"
	"github.com/foxiswho/shop-go/module/cache"
	"github.com/foxiswho/shop-go/router/base"
	"github.com/foxiswho/shop-go/router/example/api"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

//-----
// API RoutersApi
//-----
func RoutersApi() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Context自定义
	e.Use(base.NewBaseContext())

	// Customization
	if Conf.ReleaseMode {
		e.Debug = false
	}
	e.Logger.SetPrefix("api")
	e.Logger.SetLevel(GetLogLvl())

	// Session
	//e.Use(session.Session())

	// OpenTracing
	if !Conf.Opentracing.Disable {
		e.Use(opentracing.OpenTracing("api"))
	}

	// CSRF
	//e.Use(mw.CSRFWithConfig(mw.CSRFConfig{
	//	TokenLookup: "form:X-XSRF-TOKEN",
	//}))

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
	// e.GET("/user_service/:id", ec.CachePage(ec.NewMemcachedStore([]string{conf.MEMCACHED_SERVER}, time.Hour), time.Minute, UserHandler))

	// RoutersApi
	//e.GET("/login", UserLoginHandler)
	//e.GET("/register", UserRegisterHandler)

	// Unauthenticated route
	e.GET("/", accessible)
	// JWT
	j := e.Group("/jwt")
	{
		// Login route
		j.POST("/jwt-login", base.Handler(api.JwtLoginPostHandler))
		i:=j.Group("/restricted")
		{
			// Configure middleware with the custom claims type
			config := mw.JWTConfig{
				Claims:     &api.JwtCustomClaims{},
				SigningKey: []byte(Conf.SessionSecretKey),
			}
			i.Use(mw.JWTWithConfig(config))
			i.GET("/restricted", restricted)
			//j.Use(mw.JWTWithConfig(mw.JWTConfig{
			//	SigningKey:  []byte(Conf.SessionSecretKey),
			//	//ContextKey:  "_user",
			//	//TokenLookup: "header:" + echo.HeaderAuthorization,
			//}))
			i.GET("/xx", api.JwtApiHandler)
		}


		//curl http://echo.api.localhost:8080/restricted/user -H "Authorization: Bearer XXX"
		//r.GET("/user_service", UserHandler)
	}

	//post := r.Group("/post")
	//{
	//	post.GET("/save", PostSaveHandler)
	//	post.GET("/id/:id", PostHandler)
	//	post.GET("/:userId/p/:p/s/:s", PostsHandler)
	//}

	return e
}
func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*api.JwtCustomClaims)
	name := claims.Name
	c.Response().Header().Del("Access-Control-Allow-Origin")
	c.Response().Header().Add("Access-Control-Allow-Origin","*")
	return c.String(http.StatusOK, "Welcome "+name+"!")
}