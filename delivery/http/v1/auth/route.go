package auth

import (
	"github.com/auth_service/usecase"
	"github.com/auth_service/usecase/auth"
	"github.com/labstack/echo/v4"
)

type Route struct {
	authUseCase auth.IUsecase
}

func Init(group *echo.Group, useCase *usecase.UseCase) {
	r := &Route{
		authUseCase: useCase.Auth,
	}

	group.POST("/register", r.Register)
	group.POST("/login", r.Login)
}
