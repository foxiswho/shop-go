package auth

import (
	"github.com/casbin/casbin"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	m_ac "github.com/foxiswho/shop-go/module/auth/casbin"
	"strconv"
	"github.com/foxiswho/shop-go/module/log"
)

type (
	// Config defines the config for CasbinAuth middleware.
	Config struct {
		// Skipper defines a function to skip middleware.
		Skipper middleware.Skipper

		// Enforcer CasbinAuth main rule.
		// Required.
		Enforcer *casbin.Enforcer
	}
)

var (
	// DefaultConfig is the default CasbinAuth middleware config.
	DefaultConfig = Config{
		Skipper: middleware.DefaultSkipper,
	}
)

// Middleware returns a CasbinAuth middleware.
//
// For valid credentials it calls the next handler.
// For missing or invalid credentials, it sends "401 - Unauthorized" response.
func Middleware(ce *casbin.Enforcer) echo.MiddlewareFunc {
	c := DefaultConfig
	c.Enforcer = ce
	return MiddlewareWithConfig(c)
}

// MiddlewareWithConfig returns a CasbinAuth middleware with config.
// See `Middleware()`.
func MiddlewareWithConfig(config Config) echo.MiddlewareFunc {
	// Defaults
	if config.Skipper == nil {
		config.Skipper = DefaultConfig.Skipper
	}
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) || config.CheckPermission(c) {
				return next(c)
			}
			return echo.ErrForbidden
		}
	}
}

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func (a *Config) CheckPermission(c echo.Context) bool {
	role_id := m_ac.GetRoleId(c)
	method := c.Request().Method
	path := c.Request().URL.Path
	role := strconv.Itoa(role_id)
	log.Debugf("role_id ? ,method ?, ?", role, path, method)
	if a.Enforcer.Enforce(role, path, method) {
		return true
	}
	// 更多 角色权限判断
	more := m_ac.GetRoleExtend(c)
	if more != nil {
		if len(more) > 0 {
			for _, v := range more {
				if v > 0 {
					role = strconv.Itoa(v)
					log.Debugf("role_id ? ,method ?, ?", role, path, method)
					if a.Enforcer.Enforce(role, path, method) {
						return true
					}
				}
			}
		}
	}
	return false
}
