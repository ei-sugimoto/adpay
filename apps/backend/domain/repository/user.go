package repository

import (
	"context"

	"github.com/ei-sugimoto/adpay/apps/backend/domain/entity"
)

type UserRepository interface {
	Save(ctx context.Context, user entity.User) error
	ExistByName(ctx context.Context, user entity.User) (bool, error)
	ExistByID(ctx context.Context, user entity.User) (bool, error)
	GetByName(ctx context.Context, user entity.User) (entity.User, error)
}
