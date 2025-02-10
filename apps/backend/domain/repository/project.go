package repository

import (
	"context"

	"github.com/ei-sugimoto/adpay/apps/backend/domain/entity"
)

type ProjectRepository interface {
	Save(ctx context.Context, project entity.Project) error
}
