package jwt

import (
	"github.com/labstack/echo/middleware"
	jwt2 "github.com/foxiswho/shop-go/consts/session/jwt"
	"github.com/labstack/echo"
)

func GetJwtMiddlewareUser() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(GetJwtMiddleware(jwt2.Jwt_Context_Key_admin))
}

func GetJwtMiddlewareAdmin() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(GetJwtMiddleware(jwt2.Jwt_Context_Key_admin))
}
