package task

import (
	"context"
	"errors"
	"task-service/internal/entity"
	"task-service/internal/repository/task"
	"task-service/internal/repository/user"
)

var (
	ErrUserNotFound = errors.New("user was not found")
)

// Service interface
type taskService struct {
	taskRepo task.Repo
	userRepo user.Repo
}

func NewService(tr task.Repo, ur user.Repo) UseCase {
	return &taskService{
		taskRepo: tr,
		userRepo: ur,
	}
}

func (t *taskService) GetTasks(ctx context.Context) ([]*entity.Task, error) {
	return t.taskRepo.GetTasks(ctx)
}

func (t *taskService) Create(ctx context.Context, name, userID string) (*entity.Task, error) {
	usr, err := t.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if usr == nil {
		return nil, ErrUserNotFound
	}
	return t.taskRepo.Create(ctx, name, userID)
}

func (t *taskService) GetByID(ctx context.Context, ID string) (*entity.Task, error) {
	return t.taskRepo.GetByID(ctx, ID)
}

func (t *taskService) Delete(ctx context.Context, ID string) error {
	return t.taskRepo.Delete(ctx, ID)
}

func (t *taskService) Update(ctx context.Context, ID string, Name string) (*entity.Task, error) {
	return t.taskRepo.Update(ctx, ID, Name)
}
