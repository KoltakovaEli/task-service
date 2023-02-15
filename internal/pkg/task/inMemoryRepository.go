package task

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

type inMemoryRepository struct {
	tasks []*Task
}

func NewInMemoryRepository() *inMemoryRepository {
	tasks := []*Task{}
	return &inMemoryRepository{tasks}
}

func (r *inMemoryRepository) GetTasks(ctx context.Context) []*Task {
	return r.tasks
}

func (r *inMemoryRepository) Create(ctx context.Context, name string) Task {
	task := Task{
		ID:   uuid.NewString(),
		Name: name,
	}
	r.tasks = append(r.tasks, &task)
	return task
}

func (r *inMemoryRepository) GetByID(ctx context.Context, ID string) (Task, error) {
	for _, t := range r.tasks {
		if t.ID == ID {
			return *t, nil
		}
	}
	return Task{}, errors.New("task is not found")
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
func remove(slice []*Task, s int) []*Task {
	return append(slice[:s], slice[s+1:]...)
}

func (r *inMemoryRepository) Update(ctx context.Context, ID string, Name string) (Task, error) {
	for _, t := range r.tasks {
		if t.ID == ID {
			t.Name = Name
			return *t, nil
		}
	}
	return Task{}, errors.New("task is not found")
}
