package persistence

import (
	"github.com/ei-sugimoto/adpay/apps/backend/domain/entity"
	"github.com/ei-sugimoto/adpay/apps/backend/domain/repository"
)

type UserPersistence struct {
}

func NewUserPersistence() repository.UserRepository {
	return &UserPersistence{}
}

func (p *UserPersistence) Save(user entity.User) error {
	return nil
}
