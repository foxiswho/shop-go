package user_auth

import "github.com/labstack/echo"

func GetRoleId(c echo.Context) int {
	user := DefaultGetUser(c)
	return user.RoleId()
}
