package usecase

import (
	"context"

	"github.com/ei-sugimoto/adpay/apps/backend/domain/entity"
	"github.com/ei-sugimoto/adpay/apps/backend/domain/repository"
	"github.com/ei-sugimoto/adpay/apps/backend/infra/persistence"
	"github.com/ei-sugimoto/adpay/apps/backend/utils"
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
	if exist, _ := u.UserRepository.ExistByName(ctx, user); exist {
		return persistence.ErrExistUser
	}
	return u.UserRepository.Save(ctx, user)
}

func (u *UserUsecase) Login(ctx context.Context, name, password string) (string, error) {
	user := entity.NewUserWithoutID(name, password)
	getUser, err := u.UserRepository.GetByName(ctx, user)
	if err != nil {
		return "", err
	}

	if !getUser.Password.Verify(password) {
		return "", errors.New("name or password is incorrect")
	}

	token, err := utils.NewToken(getUser.ID.Int64())
	if err != nil {
		return "", err
	}

	return token, nil

}
