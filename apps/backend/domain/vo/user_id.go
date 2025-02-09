package vo

type ID int64

func NewID(id int64) ID {
	return ID(id)
}

func (id ID) Int64() int64 {
	return int64(id)
}

func (id ID) Compare(compareID int64) bool {
	return id.Int64() == compareID
}
