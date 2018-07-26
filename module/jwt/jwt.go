package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/foxiswho/shop-go/module/conf"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/foxiswho/shop-go/util/conv"
)

// jwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	Id   int    `json:"id"`
	Type string `json:"type"`
	jwt.StandardClaims
}

func GetJwtToken(id int, type_jwt string) (string, error) {
	// Set custom claims
	claims := &JwtCustomClaims{}
	claims.Id = id
	claims.Type = type_jwt
	claims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(conf.Conf.SessionSecretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func GetJwtMiddleware(ContextKey string) middleware.JWTConfig {
	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		SigningKey:  []byte(conf.Conf.SessionSecretKey),
		ContextKey:  ContextKey,
		TokenLookup: "header:" + echo.HeaderAuthorization,
	}
	return config
}

func GetJwtClaims(token *jwt.Token) map[string]interface{} {
	myMap := make(map[string]interface{})
	//fmt.Println(token.Claims)
	if token.Claims == nil {
		return myMap
	}
	myMap, _ = conv.ObjToMap(token.Claims)
	//for index, value := range maps {
	//	myMap[index] = value
	//}
	return myMap
}
