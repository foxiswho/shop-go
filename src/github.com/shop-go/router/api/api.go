package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	// "github.com/jinzhu/gorm"

	"github.com/hb-go/echo-web/middleware/session"

	"github.com/hb-go/echo-web/model"
	"github.com/hb-go/echo-web/module/cache"
	"github.com/hb-go/echo-web/module/log"
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
			cacheStore.Set("userId", 1, 5*time.Minute)
		} else {
			if err := cacheStore.Get("userId", &value); err != nil {
				log.Debugf("cache userId get err:%v", err)
			}
		}
	}

	// Flash测试
	s := session.Default(c)
	s.AddFlash("0")
	s.AddFlash("1")
	s.AddFlash("10", "key1")
	s.AddFlash("20", "key2")
	s.AddFlash("21", "key2")

	request := c.Request()
	c.AutoFMT(http.StatusOK, map[string]interface{}{
		"title":        "Api Index",
		"User":         u,
		"CacheValue":   value,
		"URL":          request.URL,
		"Scheme":       request.URL.Scheme,
		"Host":         request.Host,
		"UserAgent":    request.UserAgent(),
		"Method":       request.Method,
		"URI":          request.RequestURI,
		"RemoteAddr":   request.RemoteAddr,
		"Path":         request.URL.Path,
		"QueryString":  request.URL.RawQuery,
		"QueryParams":  request.URL.Query(),
		"HeaderKeys":   request.Header,
		"FlashDefault": s.Flashes(),
		"Flash1":       s.Flashes("key1"),
		"Flash2":       s.Flashes("key2"),
	})

	return nil
}

func JETTesterHandler(c echo.Context) error {
	t, err := getJETToken()
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
