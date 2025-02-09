package entity

import "github.com/ei-sugimoto/adpay/apps/backend/domain/vo"

type User struct {
	ID       vo.ID
	Name     vo.Name
	Password vo.Password
}

func NewUser(name, password string) User {
	return User{
		Name:     vo.NewName(name),
		Password: vo.NewPassword(password),
	}
}
