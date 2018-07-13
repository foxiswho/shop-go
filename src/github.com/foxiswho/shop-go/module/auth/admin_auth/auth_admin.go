// Base on https://github.com/martini-contrib/sessionauth
//这里不能随便自定义，否则，需要改的好多地方，如 模板
package admin_auth

import (
	"fmt"
	"net/http"

	"github.com/foxiswho/shop-go/middleware/session"
	"github.com/labstack/echo"
	"github.com/foxiswho/shop-go/module/auth/user_auth"
)

const (
	DefaultKey  = "github.com/foxiswho/shop-go/modules/auth/admin"
	errorFormat = "[modules] ERROR! %s\n"
)

var (
	// RedirectUrl should be the relative URL for your login route
	RedirectUrl string = "/admin_login/login"

	// RedirectParam is the query string parameter that will be set
	// with the page the user_service was trying to visit before they were
	// intercepted.
	RedirectParam string = "return_url"

	// SessionKey is the key containing the unique ID in your session
	SessionKey string = "ADMINAUTHUNIQUEID"
)

type AuthAdmin struct {
	user_auth.User
}

func New(newAdmin func() user_auth.User) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			s := session.Default(c)
			userId := s.Get(SessionKey)
			fmt.Println("admin userId =>", userId)
			user := newAdmin()
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
			fmt.Println("admin user=>", user)
			auth := AuthAdmin{user}
			c.Set(DefaultKey, auth)
			return next(c)
		}
	}
}

// shortcut to get AuthAdmin
func Default(c echo.Context) AuthAdmin {
	return c.Get(DefaultKey).(AuthAdmin)
}

// shortcut to get AuthAdmin
func DefaultGetAdmin(c echo.Context) user_auth.User {
	user := c.Get(DefaultKey).(AuthAdmin)
	return user.User
}

// AuthenticateSession will mark the session and user_service object as authenticated. Then
// the Login() user_service function will be called. This function should be called after
// you have validated a user_service.
func AuthenticateSession(s session.Session, user user_auth.User) error {
	user.Login()
	return UpdateUser(s, user)
}

func (a AuthAdmin) LogoutTest(s session.Session) {
	a.User.Logout()
	s.Delete(SessionKey)
	s.Save()
}

// Logout will clear out the session and call the Logout() user_service function.
func Logout(s session.Session, user user_auth.User) {
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
func UpdateUser(s session.Session, user user_auth.User) error {
	s.Set(SessionKey, user.UniqueId())
	s.Save()
	return nil
}
