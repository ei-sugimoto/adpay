package usecase

import (
	"context"

	"github.com/ei-sugimoto/adpay/apps/backend/domain/entity"
	"github.com/ei-sugimoto/adpay/apps/backend/domain/repository"
)

type ProjectUsecase struct {
	projectRepo repository.ProjectRepository
}

func NewProjectUsecase(projectRepo repository.ProjectRepository) *ProjectUsecase {
	return &ProjectUsecase{
		projectRepo: projectRepo,
	}
}

func (u *ProjectUsecase) Save(ctx context.Context, entity entity.Project) error {
	return u.projectRepo.Save(ctx, entity)
}
