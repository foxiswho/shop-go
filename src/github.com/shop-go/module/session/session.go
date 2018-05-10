package session

import (
	"github.com/labstack/echo"

	es "github.com/hb-go/echo-web/middleware/session"

	. "github.com/hb-go/echo-web/conf"
)

func Session() echo.MiddlewareFunc {
	switch Conf.SessionStore {
	case REDIS:
		store, err := es.NewRedisStore(10, "tcp", Conf.Redis.Server, Conf.Redis.Pwd, []byte("secret"))
		if err != nil {
			panic(err)
		}
		return es.New("mysession", store)
	case FILE:
		store := es.NewFilesystemStore("", []byte("secret-key"))
		return es.New("mysession", store)
	default:
		store := es.NewCookieStore([]byte("secret"))
		return es.New("mysession", store)
	}
}
