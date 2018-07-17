package casbin

import (
	"github.com/labstack/echo"
	"github.com/foxiswho/shop-go/module/auth"
)

func GetRoleId(c echo.Context) int {
	return auth.GetAuthDataRoleId(c)
}
