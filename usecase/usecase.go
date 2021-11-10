package usecase

import (
	"github.com/cnson19700/auth_service/usecase/auth"
	"github.com/cnson19700/user_service/repository"
)

type UseCase struct {
	Auth auth.IUsecase
}

func New(repo *repository.Repository) *UseCase {
	return &UseCase{

		Auth: auth.New(repo),
	}
}
