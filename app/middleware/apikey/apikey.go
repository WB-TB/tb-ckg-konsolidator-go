package middleware

import (
	"fhir-sirs/app/config"
	"net/http"

	"github.com/labstack/echo/v4"
)

func APIKeyAuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			key := c.Request().Header.Get("X-API-Key")
			if key != config.GetConfig().APIKeySecret {
				return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Unauthorized"})
			}
			return next(c)
		}
	}
}
