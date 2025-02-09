package repository

import (
	"context"

	"github.com/ei-sugimoto/adpay/apps/backend/domain/entity"
)

type UserRepository interface {
	Save(ctx context.Context, user entity.User) error
}
