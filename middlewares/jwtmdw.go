package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/cnson19700/auth_service/config"
	"github.com/cnson19700/auth_service/model"
)

type JwtCustomClaims struct {
	UserInfo UserInfo `json:"user_info"`
	jwt.StandardClaims
}

type UserInfo struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func NewClaims(user *model.User) *JwtCustomClaims {
	return &JwtCustomClaims{
		UserInfo{
			user.ID,
			user.Email,
			user.Role,
		},
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 2).Unix(),
		},
	}
}

func GetUserInfFromToken(c echo.Context) *UserInfo {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtCustomClaims)

	return &UserInfo{
		claims.UserInfo.ID,
		claims.UserInfo.Email,
		claims.UserInfo.Role,
	}
}

var Config = middleware.JWTConfig{
	Claims:     &JwtCustomClaims{},
	SigningKey: []byte(config.GetConfig().Jwt.Key),
}
