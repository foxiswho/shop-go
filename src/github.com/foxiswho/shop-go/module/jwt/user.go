package jwt

import (
	"github.com/labstack/echo/middleware"
	jwt2 "github.com/foxiswho/shop-go/consts/session/jwt"
	"github.com/labstack/echo"
)

func GetJwtMiddlewareAdminConfig() middleware.JWTConfig {
	return GetJwtMiddleware(jwt2.ContextKey_admin)
}

func GetJwtMiddlewareAdmin() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(GetJwtMiddleware(jwt2.ContextKey_admin))
}
