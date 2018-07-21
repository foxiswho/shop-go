package jwt

import (
	mw "github.com/labstack/echo/middleware"
)

func GetJwtMiddlewareAdmin() mw.JWTConfig {
	return GetJwtMiddleware(ContextKey_admin)
}
