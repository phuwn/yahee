package middleware

import (
	"strings"

	"github.com/labstack/echo"
	"github.com/phuwn/tools/errors"

	"github.com/phuwn/yahee/src/model"
)

var noAuthPath = map[string]bool{"/healthz": true, "/auth": true}

func authenticate(c echo.Context) error {
	auth := c.Request().Header.Get("Authorization")
	if !strings.Contains(auth, "Bearer ") {
		return errors.New("invalid auth method", 401)
	}
	token := auth[7:]
	if token == "" {
		return errors.New("missing access_token", 401)
	}
	uid, err := model.VerifyUserSession(token)
	if err != nil {
		return err
	}
	model.SetUserIDToCtx(c, uid)
	return nil
}

// WithAuth - authentication middleware
func WithAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if noAuthPath[c.Request().RequestURI] {
			return next(c)
		}
		err := authenticate(c)
		if err != nil {
			return errors.Customize(err, 401, "invalid token")
		}
		return next(c)
	}
}
