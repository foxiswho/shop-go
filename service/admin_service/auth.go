package admin_service

import (
	"github.com/foxiswho/shop-go/service/admin_service/auth"
	"github.com/foxiswho/shop-go/util"
	"github.com/foxiswho/shop-go/consts/session/jwt"
	jwt2 "github.com/foxiswho/shop-go/module/jwt"
)

func Login(username, password string) (string, error) {
	a := auth.NewAdminAuth()
	a.SetUsername(username)
	a.SetPassword(password)
	a.SetIsVerificationCode(false)
	admin, err := a.Process()
	if err != nil {
		return "", err
	}
	if admin != nil {
		token, err := jwt2.GetJwtToken(admin.Id, jwt.TYPE_ADMIN)
		if err != nil {
			return "", err
		}
		return token, nil
	}
	return "", util.NewError("数据错误")
}
