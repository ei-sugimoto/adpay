package persistence

import (
	"context"

	"github.com/ei-sugimoto/adpay/apps/backend/domain/entity"
	"github.com/ei-sugimoto/adpay/apps/backend/domain/repository"
	"github.com/uptrace/bun"
)

type ProjectPersistence struct {
	UserPersistence repository.UserRepository
	DB              *bun.DB
}

type Project struct {
	ID       int64  `bun:"id,pk,autoincrement"`
	Name     string `bun:"name"`
	AuthorID int64  `bun:"author_id"`
}

func NewProjectPersistence(db *bun.DB, userPersistence repository.UserRepository) repository.ProjectRepository {
	return &ProjectPersistence{UserPersistence: userPersistence, DB: db}
}

func (p *ProjectPersistence) Save(ctx context.Context, project entity.Project) error {
	var err error
	isExsist, err := p.UserPersistence.ExistByID(ctx, project.AuthorID)
	if err != nil {
		return err
	}

	if !isExsist {
		return ErrNoExistUser
	}

	convertProject := ConvertProject(project)
	_, err = p.DB.NewInsert().Model(&convertProject).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func ConvertProject(project entity.Project) Project {
	return Project{
		ID:       project.AuthorID.Int64(),
		Name:     project.Name.String(),
		AuthorID: project.AuthorID.Int64(),
	}
}
