package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	JWTConfig struct {
		Skipper    Skipper
		SigningKey interface{}
	}
	Skipper      func(e echo.Context) bool
	jwtExtractor func(echo.Context) (string, error)
)

var (
	ErrJWTMissing = echo.NewHTTPError(http.StatusUnauthorized, "missing or malformed jwt")
	ErrJWTExpired = echo.NewHTTPError(http.StatusForbidden, "invalid or expired jwt")
)

func JWT(key interface{}) echo.MiddlewareFunc {
	cfg := JWTConfig{}
	cfg.SigningKey = key
	return JWTConfig(cfg)
}
