package persistence

import (
	"context"

	"github.com/ei-sugimoto/adpay/apps/backend/domain/entity"
	"github.com/ei-sugimoto/adpay/apps/backend/domain/repository"
	"github.com/ei-sugimoto/adpay/apps/backend/domain/vo"
	"github.com/uptrace/bun"
)

type UserPersistence struct {
	DB *bun.DB
}

type User struct {
	ID       int64  `bun:"id,pk,autoincrement"`
	Name     string `bun:"name"`
	Password string `bun:"password"`
}

func NewUserPersistence(db *bun.DB) repository.UserRepository {
	return &UserPersistence{
		DB: db,
	}
}

func (p *UserPersistence) Save(ctx context.Context, user entity.User) error {
	cryptoPassword := user.Password.Crypto()
	user.Password = cryptoPassword
	convertUser := ConvertUser(user)
	_, err := p.DB.NewInsert().Model(&convertUser).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (p *UserPersistence) GetByName(ctx context.Context, user entity.User) (entity.User, error) {

	convertUser := ConvertUser(user)
	err := p.DB.NewSelect().Model(&convertUser).Where("name = ?", user.Name).Scan(ctx)
	if err != nil {
		return entity.User{}, err
	}
	return entity.User{
		ID:       vo.NewID(convertUser.ID),
		Name:     vo.NewName(convertUser.Name),
		Password: vo.NewPassword(convertUser.Password),
	}, nil
}

func ConvertUser(user entity.User) User {
	return User{
		ID:       user.ID.Int64(),
		Name:     user.Name.String(),
		Password: user.Password.String(),
	}
}
