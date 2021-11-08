package repository

import (
	"context"

	"github.com/Auth-Service/repository/user"
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
