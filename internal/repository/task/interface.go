package task

import (
	"context"
	"task-service/internal/entity"
)

type Repo interface {
	GetTasks(ctx context.Context) ([]*entity.Task, error)
	Create(ctx context.Context, name, userID string) (*entity.Task, error)
	GetByID(ctx context.Context, ID string) (*entity.Task, error)
	Delete(ctx context.Context, ID string) error
	Update(ctx context.Context, ID string, Name string) (*entity.Task, error)
}
