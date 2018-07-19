package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"github.com/foxiswho/shop-go/module/conf"
	mw "github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
)

// jwtCustomClaims are custom claims extending default ones.
type JwtCustomClaims struct {
	Id    int  `json:"id"`
	Admin bool `json:"admin"`
	jwt.StandardClaims
}

func GetJwtToken(id int) (string, error) {
	// Set custom claims
	claims := &JwtCustomClaims{
		id,
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
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

func GetJwtMiddleware() mw.JWTConfig {
	// Configure middleware with the custom claims type
	config := mw.JWTConfig{
		SigningKey:  []byte(conf.Conf.SessionSecretKey),
		ContextKey:  "_user",
		TokenLookup: "header:" + echo.HeaderAuthorization,
	}
	return config
}
