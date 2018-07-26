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
	serviceAdminAuth "github.com/foxiswho/shop-go/service/admin_service/auth"
	web_index "github.com/foxiswho/shop-go/router/web/index"
	web_test "github.com/foxiswho/shop-go/router/example/test"
	example_admin "github.com/foxiswho/shop-go/router/example/admin"
	"github.com/foxiswho/shop-go/router/example/api"
	"github.com/foxiswho/shop-go/middleware/authadapter"
	"github.com/casbin/casbin"
	auth_casbin "github.com/foxiswho/shop-go/middleware/auth"
	rbac2 "github.com/foxiswho/shop-go/router/example/admin/rbac"
	"github.com/foxiswho/shop-go/module/auth/admin_auth"
	"github.com/foxiswho/shop-go/module/context"
	"github.com/foxiswho/shop-go/module/auth/auth_middleware"
	"github.com/foxiswho/shop-go/module/jwt"
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
	////////////////////////////
	j := e.Group("/jwt")
	{
		//j.Use(context.SetSessionTypeJwt())
		j.POST("/login", context.Handler(web_test.JwtLoginPostHandler))
		i := j.Group("/restricted")
		{
			i.Use(jwt.GetJwtMiddlewareUser())
			i.GET("/xx", context.Handler(api.JwtApiHandler))
		}
	}
	////////////////////////////
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
		//
		test := index.Group("/example/test")
		{
			test.GET("/jwt/tester", context.Handler(web_test.JwtTesterHandler))
			test.GET("/jwt/login", context.Handler(web_test.JwtLoginHandler))
			//test.POST("/jwt/login", context.Handler(web_test.JwtLoginPostHandler))
			test.GET("/ws", context.Handler(web_test.WsHandler))
			test.GET("/cache", context.Handler(web_test.CacheHandler))
			test.GET("/cookie", context.Handler(web_test.NewCookie().IndexHandler))
			test.GET("/session", context.Handler(web_test.NewSession().IndexHandler))
			test.GET("/orm", context.Handler(web_test.NewOrm().IndexHandler))
			test.GET("/login", context.Handler(web_test.LoginHandler))
			test.POST("/login", context.Handler(web_test.LoginPostHandler))
			test.GET("/logout", context.Handler(web_test.LogoutHandler))
			test.GET("/register", context.Handler(web_test.RegisterHandler))
			test.POST("/register", context.Handler(web_test.RegisterPostHandler))
			user := test.Group("/user_service")
			user.Use(auth.LoginRequired())
			{
				user.GET("/:id", context.Handler(web_test.UserHandler))
			}
			test.GET("/upload", context.Handler(web_test.NewUpload().UploadIndex))
			test.POST("/upload", context.Handler(web_test.UploadPostIndex))
			test.POST("/upload-more", context.Handler(web_test.UploadMorePostIndex))
			test.POST("/upload-db", context.Handler(web_test.UploadDbHandler))
			test.GET("/jsonp", context.Handler(web_test.JsonpIndexHandler))
		}
	}
	////////////////////////////
	/////admin
	admin_login := e.Group("/admin_login")
	{
		admin_login.Use(context.SetContextTypeAdmin())
		admin_login.Use(auth_middleware.NewAdmin(serviceAdminAuth.GenerateAnonymousUser))
		admin_login.GET("/", context.Handler(example_admin.DefaultHandler))
		admin_login.GET("/login", context.Handler(example_admin.LoginHandler))
		admin_login.POST("/login", context.Handler(example_admin.LoginPostHandler))
		admin_login.GET("/logout", context.Handler(example_admin.LogoutHandler))
	}
	admin := e.Group("/admin")
	{
		admin.Use(context.SetContextTypeAdmin())
		admin_login.Use(auth_middleware.NewAdmin(serviceAdminAuth.GenerateAnonymousUser))
		admin.Use(admin_auth.LoginRequired())
		{
			admin.GET("", context.Handler(example_admin.IndexHandler))
		}
		rbac := admin.Group("/rbac")
		{
			//数据库驱动
			a := authadapter.NewAdapter("mysql", "")
			//加载 过滤条件
			ce := casbin.NewEnforcer("template/casbin/rbac_model.conf", a)
			//从数据库加载到内存中
			ce.LoadPolicy()
			//中间件
			rbac.Use(auth_casbin.Middleware(ce))
			rbac.GET("/index", context.Handler(rbac2.IndexHandler))
		}
	}
	return e
}
