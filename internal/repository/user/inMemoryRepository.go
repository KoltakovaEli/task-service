package user

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"task-service/internal/entity"
)

type inMemoryRepository struct {
	users []*entity.User
}

func NewInMemoryRepository() *inMemoryRepository {
	users := []*entity.User{}
	return &inMemoryRepository{users}
}

func (r *inMemoryRepository) GetUsers(ctx context.Context) []*entity.User {
	return r.users
}

func (r *inMemoryRepository) Create(ctx context.Context, name string) entity.User {
	user := entity.User{
		ID:    uuid.NewString(),
		Login: name,
	}
	r.users = append(r.users, &user)
	return user
}

func (r *inMemoryRepository) GetByID(ctx context.Context, ID string) (entity.User, error) {
	for _, t := range r.users {
		if t.ID == ID {
			return *t, nil
		}
	}
	return entity.User{}, errors.New("user is not found")
}

func (r *inMemoryRepository) Delete(ctx context.Context, ID string) error {
	for i, t := range r.users {
		if t.ID == ID {
			r.users = remove(r.users, i)
			return nil
		}
	}
	return errors.New("user is not exist")
}
func remove(slice []*entity.User, s int) []*entity.User {
	return append(slice[:s], slice[s+1:]...)
}

func (r *inMemoryRepository) Update(ctx context.Context, ID string, Name string) (entity.User, error) {
	for _, t := range r.users {
		if t.ID == ID {
			t.Login = Name
			return *t, nil
		}
	}
	return entity.User{}, errors.New("user is not found")
}
