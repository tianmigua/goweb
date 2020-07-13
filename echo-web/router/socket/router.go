package socket

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"letinvr.com/echo-web/model"
	"letinvr.com/echo-web/module/auth"
	"letinvr.com/echo-web/module/cache"
	"letinvr.com/echo-web/module/render"
	"letinvr.com/echo-web/module/session"
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
	e.Use(auth.New(model.GenerateAnonymousUser))

	e.GET("/ws", socketHandler)

	return e
}
