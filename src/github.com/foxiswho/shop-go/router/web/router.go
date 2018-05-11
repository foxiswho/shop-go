package web

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"

	"github.com/foxiswho/shop-go/middleware/captcha"
	"github.com/foxiswho/shop-go/middleware/staticbin"

	"github.com/foxiswho/shop-go/assets"
	. "github.com/foxiswho/shop-go/conf"
	"github.com/foxiswho/shop-go/middleware/opentracing"
	"github.com/foxiswho/shop-go/module/auth"
	"github.com/foxiswho/shop-go/module/cache"
	"github.com/foxiswho/shop-go/module/render"
	"github.com/foxiswho/shop-go/module/session"
	sauth "github.com/foxiswho/shop-go/service/user_service/auth"
	web_user "github.com/foxiswho/shop-go/router/web/user"
	web_index "github.com/foxiswho/shop-go/router/web/index"
	web_test "github.com/foxiswho/shop-go/router/web/test"
)

//---------
// Website Routers
//---------
func Routers() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Context自定义
	e.Use(NewBaseContext())
	// Customization
	if Conf.ReleaseMode {
		e.Debug = false
	}
	e.Logger.SetPrefix("web")
	e.Logger.SetLevel(GetLogLvl())

	// Session
	e.Use(session.Session())

	// CSRF
	e.Use(mw.CSRFWithConfig(mw.CSRFConfig{
		ContextKey:  "_csrf",
		TokenLookup: "form:_csrf",
	}))

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	// 验证码，优先于静态资源
	e.Use(captcha.Captcha(captcha.Config{
		CaptchaPath: "/captcha/",
		SkipLogging: true,
	}))

	// 静态资源
	switch Conf.Static.Type {
	case BINDATA:
		e.Use(staticbin.Static(assets.Asset, staticbin.Options{
			Dir:         "/",
			SkipLogging: true,
		}))
	default:
		e.Static("/assets", "./assets")
	}

	// Gzip，在验证码、静态资源之后
	// 验证码、静态资源使用http.ServeContent()，与Gzip有冲突，Nginx报错，验证码无法访问
	e.Use(mw.GzipWithConfig(mw.GzipConfig{
		Level: 5,
	}))

	// OpenTracing
	if !Conf.Opentracing.Disable {
		e.Use(opentracing.OpenTracing("web"))
	}

	// 模板
	e.Renderer = render.LoadTemplates()
	e.Use(render.Render())

	// Cache
	e.Use(cache.Cache())

	// Auth
	//e.Use(auth.New(model.GenerateAnonymousUser))
	e.Use(auth.New(sauth.GenerateAnonymousUser))
	// Routers
	e.GET("/", handler(web_index.HomeHandler))
	e.GET("/login", handler(web_user.LoginHandler))
	e.GET("/register", handler(web_user.RegisterHandler))
	e.GET("/logout", handler(web_user.LogoutHandler))
	e.POST("/login", handler(web_user.LoginPostHandler))
	e.POST("/register", handler(web_user.RegisterPostHandler))

	e.GET("/jwt/tester", handler(web_test.JWTTesterHandler))
	e.GET("/ws", handler(web_test.WsHandler))

	user := e.Group("/user_service")
	user.Use(auth.LoginRequired())
	{
		user.GET("/:id", handler(web_user.UserHandler))
	}

	about := e.Group("/about")
	about.Use(auth.LoginRequired())
	{
		about.GET("", handler(web_index.AboutHandler))
	}

	return e
}

type (
	HandlerFunc func(*BaseContext) error
)

/**
 * 自定义Context的Handler
 */
func handler(h HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(*BaseContext)
		return h(ctx)
	}
}
