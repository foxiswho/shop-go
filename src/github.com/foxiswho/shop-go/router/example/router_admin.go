package example

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"

	"github.com/foxiswho/shop-go/middleware/captcha"
	"github.com/foxiswho/shop-go/middleware/staticbin"

	"github.com/foxiswho/shop-go/assets"
	. "github.com/foxiswho/shop-go/conf"
	auth_casbin "github.com/foxiswho/shop-go/middleware/auth"
	"github.com/foxiswho/shop-go/middleware/opentracing"
	"github.com/foxiswho/shop-go/module/auth"
	"github.com/foxiswho/shop-go/module/cache"
	"github.com/foxiswho/shop-go/module/render"
	"github.com/foxiswho/shop-go/module/session"
	serviceAdminAuth "github.com/foxiswho/shop-go/service/admin_service/auth"
	example_admin "github.com/foxiswho/shop-go/router/example/admin"
	"github.com/foxiswho/shop-go/router/base"
	"github.com/foxiswho/shop-go/router/web/design"
	"github.com/casbin/casbin"
	rbac2 "github.com/foxiswho/shop-go/router/example/admin/rbac"
)

//---------
// Website Routers
//---------
func RoutersAdmin() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Context自定义
	e.Use(base.NewBaseContext())
	// Customization
	if Conf.ReleaseMode {
		e.Debug = false
	}
	e.Logger.SetPrefix("admin")
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
		e.Use(opentracing.OpenTracing("admin"))
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
	e.Use(auth.New(serviceAdminAuth.GenerateAnonymousUser))
	e.GET("/", base.Handler(example_admin.DefaultHandler))
	e.GET("/login", base.Handler(example_admin.LoginHandler))
	e.POST("/login", base.Handler(example_admin.LoginPostHandler))
	////////////////////////////
	/////admin
	admin := e.Group("/admin")
	{
		admin.GET("", base.Handler(example_admin.IndexHandler))
		des := admin.Group("/design")
		{
			des.GET("/service", base.Handler(design.ServiceMakeHandler))
		}
		rbac := admin.Group("/rbac")
		{
			ce := casbin.NewEnforcer("template/casbin/rbac_model.conf")
			rbac.Use(auth_casbin.Middleware(ce))
			rbac.GET("/index", base.Handler(rbac2.IndexHandler))
		}
	}
	// Auth
	//e.Use(auth.New(model.GenerateAnonymousUser))
	//e.Use(auth.New(serviceAuth.GenerateAnonymousUser))
	// Routers
	return e
}
