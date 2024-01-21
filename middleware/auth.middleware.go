package middleware

import (
	"ecommerce/api"
	"ecommerce/utils"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// AuthMiddleware verifies the JWT token for protected routes
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return api.Response(c, http.StatusUnauthorized, api.ResponseType{Error: errors.New("Unauthorized")})
		}

		_, err := utils.VerifyAccessToken(token)
		if err != nil {
			return api.Response(c, http.StatusUnauthorized, api.ResponseType{Error: errors.New("Unauthorized")})
		}

		// You can access user claims like claims.Subject, claims.ExpiresAt, etc.

		return next(c)
	}
}
