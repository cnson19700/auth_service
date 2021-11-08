package middlewares

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func CheckAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JwtCustomClaims)
		role := claims.UserInfo.Role

		if role == "admin" {
			return next(c)
		}

		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Administrator Only !!!",
		})
	}
}
