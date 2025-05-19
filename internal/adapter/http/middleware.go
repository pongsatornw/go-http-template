package http

import (
	"errors"
	"net/http"
	"pongsatorn/go-http-template/internal/adapter/auth/token"
	"pongsatorn/go-http-template/internal/core/domain"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const bearerPrefix = "Bearer "

type CustomMiddleware struct {
	jwtManager *token.JwtManager
}

func NewCustomMiddleware(jwtManager *token.JwtManager) *CustomMiddleware {
	return &CustomMiddleware{
		jwtManager: jwtManager,
	}
}

func (mw *CustomMiddleware) authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, domain.ErrNoAuthProvided)
		}

		if !strings.HasPrefix(authHeader, bearerPrefix) {
			return echo.NewHTTPError(http.StatusUnauthorized, domain.ErrInvalidAuthProvided)
		}

		tokenStr := strings.TrimPrefix(authHeader, bearerPrefix)
		_, err := mw.jwtManager.Validate(c.Request().Context(), tokenStr)
		if err != nil {
			if errors.Is(err, domain.ErrTokenExpired) {
				return echo.NewHTTPError(http.StatusUnauthorized, err)
			}

			return echo.NewHTTPError(http.StatusUnauthorized, domain.ErrInvalidAuthProvided)
		}

		return next(c)
	}
}

func (mw *CustomMiddleware) loggingMiddleware() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}", "method":"${method}", "uri":"${uri}",` +
			`"status":${status},c"latency_human":"${latency_human}"` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	})
}
