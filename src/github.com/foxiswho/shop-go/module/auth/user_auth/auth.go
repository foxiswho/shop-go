// Base on https://github.com/martini-contrib/sessionauth
package user_auth

import (
	"fmt"
	"net/http"

	"github.com/foxiswho/shop-go/middleware/session"
	"github.com/labstack/echo"
	"github.com/foxiswho/shop-go/module/context"
	"github.com/foxiswho/shop-go/module/context/session_type"
)

const (
	DefaultKey  = "github.com/foxiswho/shop-go/modules/auth/user_auth"
	errorFormat = "[modules] ERROR! %s\n"
)

var (
	// RedirectUrl should be the relative URL for your login route
	RedirectUrl string = "/login"

	// RedirectParam is the query string parameter that will be set
	// with the page the user_service was trying to visit before they were
	// intercepted.
	RedirectParam string = "return_url"

	// SessionKey is the key containing the unique ID in your session
	SessionKey string = "AUTHUNIQUEID"
)

type User interface {
	// Return whether this user_service is logged in or not
	IsAuthenticated() bool

	// Set any flags or extra data that should be available
	Login()

	// Clear any sensitive data out of the user_service
	Logout()

	// Return the unique identifier of this user_service object
	UniqueId() interface{}

	RoleId() int

	//MORE ROLE ID
	RoleExtend() []int

	// Populate this user_service object with values
	GetById(id interface{}) error

	Module() string
}

type AuthUser struct {
	User
}

func New(newUser func() User) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c context.BaseContext) error {
			userId := session_type.GetUserId(c)
			user := newUser()
			if userId > 0 {
				err := user.GetById(userId)
				if err != nil {
					c.Logger().Errorf("Login Error: %v", err)
				} else {
					user.Login()
				}
			} else {
				c.Logger().Debugf("Login status: No UserId")
			}

			auth := AuthUser{user}
			c.Set(DefaultKey, auth)
			return next(c)
		}
	}
}

// shortcut to get AuthUser
func Default(c echo.Context) AuthUser {
	// return c.MustGet(DefaultKey).(auth)
	return c.Get(DefaultKey).(AuthUser)
}

// shortcut to get AuthUser
func DefaultGetUser(c echo.Context) User {
	// return c.MustGet(DefaultKey).(auth)
	auth := c.Get(DefaultKey).(AuthUser)
	return auth.User
}

// AuthenticateSession will mark the session and user_service object as authenticated. Then
// the Login() user_service function will be called. This function should be called after
// you have validated a user_service.
func AuthenticateSession(s session.Session, user User) error {
	user.Login()
	return UpdateUser(s, user)
}

func (a AuthUser) LogoutTest(s session.Session) {
	a.User.Logout()
	s.Delete(SessionKey)
	s.Save()
}

// Logout will clear out the session and call the Logout() user_service function.
func Logout(s session.Session, user User) {
	user.Logout()
	s.Delete(SessionKey)
	s.Save()
}

// LoginRequired verifies that the current user_service is authenticated. Any routes that
// require a login should have this handler placed in the flow. If the user_service is not
// authenticated, they will be redirected to /login with the "next" get parameter
// set to the attempted URL.
func LoginRequired() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			a := Default(c)
			if a.User.IsAuthenticated() == false {
				uri := c.Request().RequestURI
				path := fmt.Sprintf("%s?%s=%s", RedirectUrl, RedirectParam, uri)
				c.Redirect(http.StatusMovedPermanently, path)
				return echo.NewHTTPError(http.StatusUnauthorized)
			}

			return next(c)
		}
	}
}

// UpdateUser updates the Admin object stored in the session. This is useful incase a change
// is made to the user_service model that needs to persist across requests.
func UpdateUser(s session.Session, user User) error {
	s.Set(SessionKey, user.UniqueId())
	s.Save()
	return nil
}
