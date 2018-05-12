package session

import (
	"github.com/labstack/echo"

	es "github.com/foxiswho/shop-go/middleware/session"

	. "github.com/foxiswho/shop-go/conf"
)

const SESSION_KEY = "SESSID"

func Session() echo.MiddlewareFunc {
	_session_key := []byte(Conf.SessionSecretKey)
	switch Conf.SessionStore {
	case REDIS:
		store, err := es.NewRedisStore(10, "tcp", Conf.Redis.Server, Conf.Redis.Pwd, _session_key)
		if err != nil {
			panic(err)
		}
		return es.New(SESSION_KEY, store)
	case FILE:
		store := es.NewFilesystemStore("", _session_key)
		return es.New(SESSION_KEY, store)
	default:
		store := es.NewCookieStore(_session_key)
		return es.New(SESSION_KEY, store)
	}
}
