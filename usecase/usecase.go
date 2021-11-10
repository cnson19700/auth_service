package usecase

import (
	"github.com/auth_service/repository"
	"github.com/auth_service/usecase/auth"
)

type UseCase struct {
	Auth auth.IUsecase
}

func New(repo *repository.Repository) *UseCase {
	return &UseCase{

		Auth: auth.New(repo),
	}
}
