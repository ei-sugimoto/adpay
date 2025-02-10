package entity

import "github.com/ei-sugimoto/adpay/apps/backend/domain/vo"

type Project struct {
	ID       vo.ProjectID
	Name     vo.ProjectName
	AuthorID vo.UserID
}

func NewProject(id int64, name string, authorID int64) Project {
	return Project{
		ID:       vo.NewProjectID(id),
		Name:     vo.NewProjectName(name),
		AuthorID: vo.NewUserID(authorID),
	}
}

func NewProjectWithoutID(name string, authorID int64) Project {
	return Project{
		Name:     vo.NewProjectName(name),
		AuthorID: vo.NewUserID(authorID),
	}
}
