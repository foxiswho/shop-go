// Base on https://github.com/martini-contrib/sessionauth
package auth

import (
	"fmt"
	"net/http"

	"github.com/hb-go/echo-web/middleware/session"
	"github.com/labstack/echo"
)

const (
	DefaultKey  = "github.com/hb-go/echo-web/modules/auth"
	errorFormat = "[modules] ERROR! %s\n"
)

var (
	// RedirectUrl should be the relative URL for your login route
	RedirectUrl string = "/login"

	// RedirectParam is the query string parameter that will be set
	// with the page the user was trying to visit before they were
	// intercepted.
	RedirectParam string = "return_url"

	// SessionKey is the key containing the unique ID in your session
	SessionKey string = "AUTHUNIQUEID"
)

type User interface {
	// Return whether this user is logged in or not
	IsAuthenticated() bool

	// Set any flags or extra data that should be available
	Login()

	// Clear any sensitive data out of the user
	Logout()

	// Return the unique identifier of this user object
	UniqueId() interface{}

	// Populate this user object with values
	GetById(id interface{}) error
}

type Auth struct {
	User
}

func New(newUser func() User) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			s := session.Default(c)
			userId := s.Get(SessionKey)
			user := newUser()

			if userId != nil {
				err := user.GetById(userId)
				if err != nil {
					c.Logger().Errorf("Login Error: %v", err)
				} else {
					user.Login()
				}
			} else {
				c.Logger().Debugf("Login status: No UserId")
			}

			auth := Auth{user}
			c.Set(DefaultKey, auth)

			return next(c)
		}
	}
}

// shortcut to get Auth
func Default(c echo.Context) Auth {
	// return c.MustGet(DefaultKey).(auth)
	return c.Get(DefaultKey).(Auth)
}

// AuthenticateSession will mark the session and user object as authenticated. Then
// the Login() user function will be called. This function should be called after
// you have validated a user.
func AuthenticateSession(s session.Session, user User) error {
	user.Login()
	return UpdateUser(s, user)
}

func (a Auth) LogoutTest(s session.Session) {
	a.User.Logout()
	s.Delete(SessionKey)
	s.Save()
}

// Logout will clear out the session and call the Logout() user function.
func Logout(s session.Session, user User) {
	user.Logout()
	s.Delete(SessionKey)
	s.Save()
}

// LoginRequired verifies that the current user is authenticated. Any routes that
// require a login should have this handler placed in the flow. If the user is not
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

// UpdateUser updates the User object stored in the session. This is useful incase a change
// is made to the user model that needs to persist across requests.
func UpdateUser(s session.Session, user User) error {
	s.Set(SessionKey, user.UniqueId())
	s.Save()
	return nil
}
