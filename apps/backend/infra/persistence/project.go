package persistence

import (
	"context"

	"github.com/ei-sugimoto/adpay/apps/backend/domain/entity"
	"github.com/ei-sugimoto/adpay/apps/backend/domain/repository"
	"github.com/uptrace/bun"
)

type ProjectPersistence struct {
	DB *bun.DB
}

type Project struct {
	ID       int64  `bun:"id,pk,autoincrement"`
	Name     string `bun:"name"`
	AuthorID int64  `bun:"author_id"`
}

func NewProjectPersistence(db *bun.DB) repository.ProjectRepository {
	return &ProjectPersistence{DB: db}
}

func (p *ProjectPersistence) Save(ctx context.Context, project entity.Project) error {
	_, err := p.DB.NewInsert().Model(&project).Exec(ctx)
	return err
}

func ConvertProject(project entity.Project) Project {
	return Project{
		ID:       project.AuthorID.Int64(),
		Name:     project.Name.String(),
		AuthorID: project.AuthorID.Int64(),
	}
}
