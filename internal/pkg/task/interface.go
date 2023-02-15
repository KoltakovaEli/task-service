package task

import "context"

type Repo interface {
	GetTasks(ctx context.Context) ([]*Task, error)
	Create(ctx context.Context, name, userID string) (Task, error)
	GetByID(ctx context.Context, ID string) (Task, error)
	Delete(ctx context.Context, ID string) error
	Update(ctx context.Context, ID string, Name string) (Task, error)
}

type Task struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	UserID string `json:"user_id"`
}
