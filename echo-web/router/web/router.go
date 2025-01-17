package web

import (
	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"

	"letinvr.com/echo-web/assets"
	. "letinvr.com/echo-web/conf"
	"letinvr.com/echo-web/middleware/captcha"
	"letinvr.com/echo-web/middleware/opentracing"
	"letinvr.com/echo-web/middleware/staticbin"
	"letinvr.com/echo-web/model"
	"letinvr.com/echo-web/module/auth"
	"letinvr.com/echo-web/module/cache"
	"letinvr.com/echo-web/module/render"
	"letinvr.com/echo-web/module/session"
)

// ---------
// Website Routers
// ---------
func Routers() *echo.Echo {
	// Echo instance
	e := echo.New()

	// Context自定义
	e.Use(NewContext())

	// Customization
	if Conf.ReleaseMode {
		e.Debug = false
	}
	e.Logger.SetPrefix("web")
	e.Logger.SetLevel(GetLogLvl())

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
	e.Static("/uploaded", "./uploaded")
	// Session
	e.Use(session.Session())

	// CSRF
	e.Use(mw.CSRFWithConfig(mw.CSRFConfig{
		ContextKey:  "_csrf",
		TokenLookup: "header:X-CSRF-Token",
	}))

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	// 验证码，优先于静态资源
	e.Use(captcha.Captcha(captcha.Config{
		CaptchaPath: "/captcha/",
		SkipLogging: true,
		StdWidth:    140,
		StdHeight:   40,
	}))

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
	e.Use(auth.New(model.GenerateAnonymousUser))

	// Routers
	e.GET("/", handler(HomeHandler))
	e.GET("/login", handler(LoginHandler))
	e.GET("/register", handler(RegisterHandler))
	e.GET("/logout", handler(LogoutHandler))
	e.POST("/login", handler(LoginPostHandler))
	e.POST("/register", handler(RegisterPostHandler))
	e.GET("/dashboard", DashboardHandler)

	e.GET("/jwt/tester", handler(JWTTesterHandler))
	e.GET("/ws", handler(WsHandler))

	user := e.Group("/user")
	user.Use(auth.LoginRequired())
	{
		user.GET("/:id", handler(UserHandler))
		user.GET("/list", handler(UListHandler))
		user.GET("/add", handler(UAddHandler))
		user.POST("/save", handler(USaveHandler))
		user.GET("/edit", handler(UEditHandler))
		user.POST("/update", handler(UUpdateHandler))
		user.POST("/delete", handler(UDeleteHandler))
	}

	about := e.Group("/about")
	about.Use(auth.LoginRequired())
	{
		about.GET("", handler(AboutHandler))
	}
	upload := e.Group("/upload")
	upload.Use(auth.LoginRequired())
	{
		upload.POST("", handler(UploadHandler))
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
