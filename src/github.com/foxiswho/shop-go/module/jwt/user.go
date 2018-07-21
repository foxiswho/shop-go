package jwt

import (
	mw "github.com/labstack/echo/middleware"
	"github.com/foxiswho/shop-go/consts/session/jwt"
)

func GetJwtMiddlewareAdmin() mw.JWTConfig {
	return GetJwtMiddleware(jwt.ContextKey_admin)
}
