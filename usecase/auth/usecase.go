package auth

import (
	"github.com/auth_service/repository"
	"github.com/auth_service/repository/user"
)

type Usecase struct {
	userRepo user.Repository
}

func New(repo *repository.Repository) IUsecase {
	return &Usecase{
		userRepo: repo.User,
	}
}
