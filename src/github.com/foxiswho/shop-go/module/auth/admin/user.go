package admin

import (
	"github.com/labstack/echo"
	"github.com/foxiswho/shop-go/module/auth/admin_auth"
)

//
func GetRoleId(c echo.Context) int {
	user := admin_auth.DefaultGetAdmin(c)
	return user.RoleId()
}

func GetRoleExtend(c echo.Context) []int {
	user := admin_auth.DefaultGetAdmin(c)
	return user.RoleExtend()
}
