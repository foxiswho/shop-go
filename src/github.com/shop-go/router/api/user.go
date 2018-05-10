package api

import (
	"net/http"
	"strconv"
	// "time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	// "github.com/jinzhu/gorm"

	"github.com/hb-go/echo-web/model"
)

func UserHandler(c echo.Context) error {
	user := c.Get("_user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	idStr := claims["id"].(string)

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		panic(err)
	}

	var User model.User
	u := User.GetUserById(id)

	c.JSON(http.StatusOK, map[string]interface{}{
		"title":  "User",
		"user":   u,
		"claims": claims,
	})

	return nil
}

func UserLoginHandler(c echo.Context) error {

	t, err := getJETToken()
	if err != nil {
		return err
	}

	c.JSON(200, map[string]interface{}{
		"URI":   "api user login",
		"token": t,
	})

	return nil
}

func UserRegisterHandler(c echo.Context) error {

	c.JSON(200, map[string]interface{}{"URI": "api user regist"})

	return nil
}

func getJETToken() (t string, e error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = "1"
	claims["name"] = "Hobo"
	claims["admin"] = true
	// claims["exp"] = time.Now().Add(time.Minute * 1).Unix()

	// Generate encoded token and send it as response.
	t, e = token.SignedString([]byte("secret"))
	return
}
