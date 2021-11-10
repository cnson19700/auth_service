package repository

import (
	"context"

	"github.com/cnson19700/auth_service/repository/user"
	"gorm.io/gorm"
)

type Repository struct {
	User user.Repository
}

func New(
	getSQLClient func(ctx context.Context) *gorm.DB,
	// getRedisClient func(ctx context.Context) *redis.Client,
) *Repository {
	return &Repository{
		User: user.NewPGRepository(getSQLClient),
	}
}
