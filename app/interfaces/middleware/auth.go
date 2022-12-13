package middleware

import (
	"flyme-backend/app/packages/auth"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Authentication(contextKey string) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(
		middleware.JWTConfig{
			ContextKey: contextKey,
			ParseTokenFunc: func(tokenStr string, c echo.Context) (interface{}, error) {
				return auth.ValidateUserToken(tokenStr)
			},
		},
	)
}
