package repository

import "github.com/ei-sugimoto/adpay/apps/backend/domain/entity"

type UserRepository interface {
	Save(user entity.User) error
}
