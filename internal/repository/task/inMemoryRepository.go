package task

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"task-service/internal/entity"
)

type inMemoryRepository struct {
	tasks []*entity.Task
}

func NewInMemoryRepository() *inMemoryRepository {
	tasks := []*entity.Task{}
	return &inMemoryRepository{tasks}
}

func (r *inMemoryRepository) GetTasks(ctx context.Context) []*entity.Task {
	return r.tasks
}

func (r *inMemoryRepository) Create(ctx context.Context, name string) entity.Task {
	task := entity.Task{
		ID:   uuid.NewString(),
		Name: name,
	}
	r.tasks = append(r.tasks, &task)
	return task
}

func (r *inMemoryRepository) GetByID(ctx context.Context, ID string) (entity.Task, error) {
	for _, t := range r.tasks {
		if t.ID == ID {
			return *t, nil
		}
	}
	return entity.Task{}, errors.New("task is not found")
}

func (r *inMemoryRepository) Delete(ctx context.Context, ID string) error {
	for i, t := range r.tasks {
		if t.ID == ID {
			r.tasks = remove(r.tasks, i)
			return nil
		}
	}
	return errors.New("task is not exist")
}
func remove(slice []*entity.Task, s int) []*entity.Task {
	return append(slice[:s], slice[s+1:]...)
}

func (r *inMemoryRepository) Update(ctx context.Context, ID string, Name string) (entity.Task, error) {
	for _, t := range r.tasks {
		if t.ID == ID {
			t.Name = Name
			return *t, nil
		}
	}
	return entity.Task{}, errors.New("task is not found")
}
