package context

import (
	"github.com/foxiswho/shop-go/consts/context"
	"github.com/foxiswho/shop-go/module/auth/admin_auth"
	"github.com/foxiswho/shop-go/module/auth/user_auth"
	"github.com/foxiswho/shop-go/middleware/session"
	"github.com/foxiswho/shop-go/util/conv"
	"github.com/foxiswho/shop-go/module/log"
)

//获取用户ID
func (c *BaseContext) GetUserId() int {
	//会话方式 jwt
	if context.Session_Type_jwt == c.SessionType {
		log.Debugf("c.ContextType == %v",c.ContextType)
		//admin 后台
		if context.Context_Type_Admin ==c.ContextType {
			claims := c.JwtTokenGetAdmin()
			if claims != nil {
				if id, ok := claims["id"]; ok {
					id2, _ := conv.ObjToInt(id)
					return id2
				}
			}
		} else if context.Context_Type_User == c.ContextType {
			//user 前台
			claims := c.JwtTokenGetUser()
			if claims != nil {
				if id, ok := claims["id"]; ok {
					id2, _ := conv.ObjToInt(id)
					return id2
				}
			}
		}
	} else if context.Session_Type_cookie == c.SessionType {
		//会话方式 cookie
		s := session.Default(c)
		//admin 后台
		if context.Context_Type_Admin ==c.ContextType {
			userId := s.Get(admin_auth.SessionKey)
			if userId != nil {
				id, _ := conv.ObjToInt(userId)
				return id
			}
		} else if context.Context_Type_User == c.ContextType {
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
