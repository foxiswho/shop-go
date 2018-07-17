package admin_auth

import (
	"github.com/labstack/echo"
)

//
func GetRoleId(c echo.Context) int {
	user := DefaultGetAdmin(c)
	return user.RoleId()
}

func GetRoleExtend(c echo.Context) []int {
	user := DefaultGetAdmin(c)
	return user.RoleExtend()
}
