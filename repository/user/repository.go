package user

import (
	"context"

	"github.com/auth_service/model"
)

type Repository interface {
	GetById(ctx context.Context, ID int64) (*model.User, error)
	GetAll(ctx context.Context) ([]model.User, error)
	Delete(ctx context.Context, ID int64) error
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
	GetEmail(ctx context.Context, email string) (*model.User, error)
	CheckEmailExist(ctx context.Context, email string) bool
	UpdatePassword(Ctx context.Context, passwordHash string, ID int64) error
}
