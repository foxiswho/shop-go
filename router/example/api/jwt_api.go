package api

import (
	"net/http"
	"time"

	"github.com/foxiswho/shop-go/middleware/session"

	"github.com/foxiswho/shop-go/module/cache"
	"github.com/foxiswho/shop-go/module/log"
	"github.com/foxiswho/shop-go/service/user_service/auth"
	userService "github.com/foxiswho/shop-go/service/example_service"
	"fmt"
	"github.com/foxiswho/shop-go/module/context"
)

func JwtApiHandler(c *context.BaseContext) error {
	//
	fmt.Println("id", c.GetUserId())
	fmt.Println("id", c.SessionType)
	id := c.GetUserId()

	u := &auth.User{}
	if id < 1 {
		log.Debugf("Render Error id: %v", id)
	} else {
		u = userService.GetUserById(uint64(id))
	}

	// 缓存测试
	value := -1
	if u != nil {
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
	c.Response().Header().Del("Access-Control-Allow-Origin")
	c.Response().Header().Add("Access-Control-Allow-Origin", "*")
	request := c.Request()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"title":        "Api Index",
		"Admin":        u,
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

}
