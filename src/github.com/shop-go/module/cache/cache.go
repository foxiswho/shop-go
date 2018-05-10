package cache

import (
	"time"

	ec "github.com/hb-go/echo-web/middleware/cache"
	"github.com/labstack/echo"

	. "github.com/hb-go/echo-web/conf"
)

const (
	DefaultExpiration = 3600
	DEFAULT           = time.Duration(0)
	FOREVER           = time.Duration(-1)
	DefaultKey        = "github.com/hb-go/echo-web/modules/cache"
)

func Cache() echo.MiddlewareFunc {
	var store ec.CacheStore

	switch Conf.CacheStore {
	case MEMCACHED:
		store = ec.NewMemcachedStore([]string{Conf.Memcached.Server}, time.Hour)
	case REDIS:
		store = ec.NewRedisCache(Conf.Redis.Server, Conf.Redis.Pwd, DefaultExpiration)
	default:
		store = ec.NewInMemoryStore(time.Hour)
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(DefaultKey, store)

			return next(c)
		}
	}
}

// shortcut to get Cache
func Default(c echo.Context) ec.CacheStore {
	// return c.MustGet(DefaultKey).(ec.CacheStore)
	return c.Get(DefaultKey).(ec.CacheStore)
}
