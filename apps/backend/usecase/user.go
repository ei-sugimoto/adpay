package usecase

import (
	"context"

	"github.com/ei-sugimoto/adpay/apps/backend/domain/entity"
	"github.com/ei-sugimoto/adpay/apps/backend/domain/repository"
	"github.com/pkg/errors"
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

func (u *UserUsecase) Login(ctx context.Context, name, password string) (entity.User, error) {
	user := entity.NewUserWithoutID(name, password)
	getUser, err := u.UserRepository.GetByName(ctx, user)
	if err != nil {
		return entity.User{}, err
	}

	if !getUser.Password.Verify(password) {
		return entity.User{}, errors.New("name or password is incorrect")
	}

	return getUser, nil
}
