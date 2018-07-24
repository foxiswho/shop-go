package context

import (
	"github.com/dgrijalva/jwt-go"
	jwt2 "github.com/foxiswho/shop-go/consts/session/jwt"
	jwt3 "github.com/foxiswho/shop-go/module/jwt"
)

func (c *BaseContext) JwtTokenGetAdmin() map[string]interface{} {
	myMap := make(map[string]interface{})
	val := c.Get(jwt2.ContextKey_admin)
	if val != nil {
		info := val.(*jwt.Token)
		if info != nil {
			return jwt3.GetJwtClaims(info)
		}
	}
	return myMap
}
