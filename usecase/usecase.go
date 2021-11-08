package usecase

import (
	"github.com/Auth-Service/repository"
	"github.com/Auth-Service/usecase/auth"
)

type UseCase struct {
	Auth auth.IUsecase
}

func New(repo *repository.Repository) *UseCase {
	return &UseCase{

		Auth: auth.New(repo),
	}
}
