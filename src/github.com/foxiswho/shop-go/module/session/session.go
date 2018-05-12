package session

import (
	"github.com/labstack/echo"

	es "github.com/foxiswho/shop-go/middleware/session"

	. "github.com/foxiswho/shop-go/conf"
)

const SESSION_KEY  = "SESSID"

func Session() echo.MiddlewareFunc {
	switch Conf.SessionStore {
	case REDIS:
		store, err := es.NewRedisStore(10, "tcp", Conf.Redis.Server, Conf.Redis.Pwd, []byte("secret"))
		if err != nil {
			panic(err)
		}
		return es.New(SESSION_KEY, store)
	case FILE:
		store := es.NewFilesystemStore("", []byte("secret-key"))
		return es.New(SESSION_KEY, store)
	default:
		store := es.NewCookieStore([]byte("secret"))
		return es.New(SESSION_KEY, store)
	}
}
