package session_type

import (
	"github.com/dgrijalva/jwt-go"
	jwt2 "github.com/foxiswho/shop-go/module/jwt"
	"github.com/foxiswho/shop-go/module/context"
)

func JwtTokenGetAdmin(c *context.BaseContext) map[string]*interface{} {
	val := c.Get(jwt2.ContextKey_admin)
	if val != nil {
		info := val.(*jwt.Token)
		if info != nil {
			return jwt2.GetJwtClaims(info)
		}
	}
	return nil
}
