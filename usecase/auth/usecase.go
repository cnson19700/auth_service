package auth

import (
	"github.com/Auth-Service/repository"
	"github.com/Auth-Service/repository/user"
)

type Usecase struct {
	userRepo user.Repository
}

func New(repo *repository.Repository) IUsecase {
	return &Usecase{
		userRepo: repo.User,
	}
}
