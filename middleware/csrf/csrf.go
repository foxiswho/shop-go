package csrf

import "github.com/labstack/echo"
import "github.com/labstack/echo/middleware"

func CSRFWithConfig() echo.MiddlewareFunc {
	return middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:X-XSRF-TOKEN",
	})
}
