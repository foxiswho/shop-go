package context

import (
	"github.com/foxiswho/shop-go/module/auth/user_auth"
	"github.com/foxiswho/shop-go/module/auth/admin_auth"
	"github.com/foxiswho/shop-go/module/jwt"
	"github.com/foxiswho/shop-go/middleware/session"
	"github.com/foxiswho/shop-go/util/conv"
)

//获取用户ID
func (c *BaseContext) GetUserId() int {
	//会话方式 jwt
	if session_jwt == c.SessionType {
		//admin 后台
		if c.ContextType == jwt.ContextKey_admin {
			claims := c.JwtTokenGetAdmin()
			if claims != nil {
				return claims["id"]
			}
		} else if jwt.ContextKey_user == c.ContextType {
			//user 前台
		}
	} else if session_cookie == c.SessionType {
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

//user
func (c *BaseContext) AuthUser() user_auth.AuthUser {
	return user_auth.Default(c)
}

//admin 后台
func (c *BaseContext) AuthAdmin() admin_auth.AuthAdmin {
	return admin_auth.Default(c)
}
