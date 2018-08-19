package web

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"

	"github.com/foxiswho/shop-go/middleware/captcha"
	"github.com/foxiswho/shop-go/middleware/staticbin"

	"github.com/foxiswho/shop-go/assets"
	. "github.com/foxiswho/shop-go/module/conf"
	"github.com/foxiswho/shop-go/middleware/opentracing"
	auth "github.com/foxiswho/shop-go/module/auth/user_auth"
	"github.com/foxiswho/shop-go/module/cache"
	"github.com/foxiswho/shop-go/module/render"
	"github.com/foxiswho/shop-go/module/session"
	serviceAuth "github.com/foxiswho/shop-go/service/user_service/auth"
	web_index "github.com/foxiswho/shop-go/router/web/index"
	"github.com/foxiswho/shop-go/module/context"
	"github.com/foxiswho/shop-go/module/auth/auth_middleware"
)

//---------
// Website Routers
//---------
func Routers() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Context自定义
	e.Use(context.NewBaseContext())
	// Customization
	if Conf.ReleaseMode {
		e.Debug = false
	}
	e.Logger.SetPrefix("web")
	e.Logger.SetLevel(GetLogLvl())

	// Session
	e.Use(session.Session())

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
	// CSRF
	e.Use(mw.CSRFWithConfig(mw.CSRFConfig{
		ContextKey:  "_csrf",
		TokenLookup: "form:_csrf",
	}))

	// 模板
	e.Renderer = render.LoadTemplates()
	e.Use(render.Render())

	// Cache
	e.Use(cache.Cache())
	// AuthUser
	//e.Use(auth.New(model.GenerateAnonymousUser))
	//e.Use(auth.New(serviceAuth.GenerateAnonymousUser))
	// Routers
	index := e.Group("")
	{
		index.Use(auth_middleware.NewUser(serviceAuth.GenerateAnonymousUser))
		index.GET("/", context.Handler(web_index.HomeHandler))
		//
		about := index.Group("/about")
		about.Use(auth.LoginRequired())
		{
			about.GET("", context.Handler(web_index.AboutHandler))
		}
	}
	return e
}
