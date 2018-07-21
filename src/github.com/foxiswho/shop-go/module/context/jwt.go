package context

import (
	"github.com/dgrijalva/jwt-go"
	jwt2 "github.com/foxiswho/shop-go/module/jwt"
)

func (ctx *BaseContext) JwtTokenGetAdmin() map[string]*interface{} {
	val := ctx.Get(jwt2.ContextKey_admin)
	if val != nil {
		info := val.(*jwt.Token)
		if info != nil {
			return jwt2.GetJwtClaims(info)
		}
	}
	return nil
}
