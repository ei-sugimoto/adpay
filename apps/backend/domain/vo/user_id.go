package vo

type UserID int64

func NewUserID(id int64) UserID {
	return UserID(id)
}

func (id UserID) Int64() int64 {
	return int64(id)
}
