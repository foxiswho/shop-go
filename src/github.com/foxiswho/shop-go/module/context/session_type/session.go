package session_type

import (
	"github.com/foxiswho/shop-go/consts/context"
	"github.com/foxiswho/shop-go/module/jwt"
	"github.com/foxiswho/shop-go/module/auth/admin_auth"
	"github.com/foxiswho/shop-go/module/auth/user_auth"
	context2 "github.com/foxiswho/shop-go/module/context"
	"github.com/foxiswho/shop-go/middleware/session"
	"github.com/foxiswho/shop-go/util/conv"
)

//获取用户ID
func GetUserId(c *context2.BaseContext) int {
	//会话方式 jwt
	if context.Session_jwt == c.SessionType {
		//admin 后台
		if c.ContextType == jwt.ContextKey_admin {
			claims := JwtTokenGetAdmin(c)
			if claims != nil {
				return claims["id"]
			}
		} else if jwt.ContextKey_user == c.ContextType {
			//user 前台
		}
	} else if context.Session_cookie == c.SessionType {
		//会话方式 cookie
		s := session.Default(c)
		//admin 后台
		if c.ContextType == jwt.ContextKey_admin {
			userId := s.Get(admin_auth.SessionKey)
			if userId != nil {
				id, _ := conv.ObjToInt(userId)
				return id
			}
		} else if jwt.ContextKey_user == c.ContextType {
			// user 前台
			userId := s.Get(user_auth.SessionKey)
			if userId != nil {
				id, _ := conv.ObjToInt(userId)
				return id
			}
		}
	}
	return 0
}
