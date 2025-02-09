package usecase

import (
	"context"

	"github.com/ei-sugimoto/adpay/apps/backend/domain/entity"
	"github.com/ei-sugimoto/adpay/apps/backend/domain/repository"
)

type UserUsecase struct {
	UserRepository repository.UserRepository
}

func NewUserUsecase(userRepository repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		UserRepository: userRepository,
	}
}

func (u *UserUsecase) Save(ctx context.Context, user entity.User) error {
	return u.UserRepository.Save(ctx, user)
}
