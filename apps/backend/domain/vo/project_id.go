package vo

type ProjectID int64

func NewProjectID(id int64) ProjectID {
	return ProjectID(id)
}
