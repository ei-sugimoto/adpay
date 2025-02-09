package entity

import "github.com/ei-sugimoto/adpay/apps/backend/domain/vo"

type User struct {
	ID       vo.UserID
	Name     vo.Name
	Password vo.Password
}

func NewUser(id int64, name, password string) User {
	return User{
		ID:       vo.NewUserID(id),
		Name:     vo.NewName(name),
		Password: vo.NewPassword(password),
	}
}

func NewUserWithoutID(name, password string) User {
	return User{
		Name:     vo.NewName(name),
		Password: vo.NewPassword(password),
	}
}
