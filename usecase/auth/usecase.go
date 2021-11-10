package auth

import (
	"github.com/cnson19700/user_service/repository"
	"github.com/cnson19700/user_service/repository/user"
)

type Usecase struct {
	userRepo user.Repository
}

func New(repo *repository.Repository) IUsecase {
	return &Usecase{
		userRepo: repo.User,
	}
}
