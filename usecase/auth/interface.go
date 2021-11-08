package auth

import (
	"context"

	"github.com/Auth-Service/model"
)

type IUsecase interface {
	Register(ctx context.Context, req RegisterRequest) (*model.User, error)
	Login(ctx context.Context, req LoginRequest) (string, error)
}
