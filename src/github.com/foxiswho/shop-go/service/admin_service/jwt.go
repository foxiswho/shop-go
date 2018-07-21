package admin_service

import (
	jwt2 "github.com/foxiswho/shop-go/module/jwt"
	mw "github.com/labstack/echo/middleware"
)

func GetJwtToken(id int) (string, error) {
	// Generate encoded token and send it as response.
	t, err := jwt2.GetJwtToken(id, jwt2.TYPE_ADMIN)
	return t, err
}

func GetJwtMiddleware() mw.JWTConfig {
	return jwt2.GetJwtMiddleware(jwt2.ContextKey_admin)
}
