package auth

import (
	"errors"

	"github.com/Auth-Service/usecase/auth"
	"github.com/cnson19700/pkg/apperror"
	"github.com/cnson19700/pkg/utils"
	"github.com/labstack/echo/v4"
)

func (r *Route) Login(c echo.Context) error {
	var (
		ctx      = &utils.CustomEchoContext{Context: c}
		appError = apperror.AppError{}
		req      = auth.LoginRequest{}
	)
	if err := c.Bind(&req); err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, apperror.ErrInvalidInput(err))
	}

	token, err := r.authUseCase.Login(ctx, req)
	if err != nil {
		_ = errors.As(err, &appError)

		return utils.Response.Error(ctx, appError)
	}

	return utils.Response.Success(ctx, token)
}
