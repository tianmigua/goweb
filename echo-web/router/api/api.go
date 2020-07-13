package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	// "github.com/jinzhu/gorm"

	"letinvr.com/echo-web/model"
	"letinvr.com/echo-web/module/cache"
	"letinvr.com/echo-web/module/log"
)

func ApiHandler(c *Context) error {
	idStr := c.QueryParam("id")
	id, err := strconv.ParseUint(idStr, 10, 64)

	u := &model.User{}
	if err != nil {
		log.Debugf("Render Error: %v", err)
	} else {
		var User model.User
		u = User.GetUserById(id)
	}

	// 缓存测试
	value := -1
	if err == nil {
		cacheStore := cache.Default(c)
		if id == 1 {
			value = 0
			if err := cacheStore.Set("userId", 1, 5*time.Minute); err != nil {
				log.Errorf("cache error:%v", err)
			}
		} else {
			if err := cacheStore.Get("userId", &value); err != nil {
				log.Debugf("cache userId get err:%v", err)
			}
		}
	}

	request := c.Request()
	c.AutoFMT(http.StatusOK, map[string]interface{}{
		"title":       "Api Index",
		"User":        u,
		"CacheValue":  value,
		"URL":         request.URL,
		"Scheme":      request.URL.Scheme,
		"Host":        request.Host,
		"UserAgent":   request.UserAgent(),
		"Method":      request.Method,
		"URI":         request.RequestURI,
		"RemoteAddr":  request.RemoteAddr,
		"Path":        request.URL.Path,
		"QueryString": request.URL.RawQuery,
		"QueryParams": request.URL.Query(),
		"HeaderKeys":  request.Header,
	})

	return nil
}

func JWTTesterHandler(c echo.Context) error {
	t, err := getJWTToken()
	if err != nil {
		return err
	}
	c.Set("tmpl", "api/jwt_tester")
	c.Set("data", map[string]interface{}{
		"title": "JWT 接口测试",
		"token": t,
	})

	return nil
}
