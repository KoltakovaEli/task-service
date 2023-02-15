package user

import (
	"context"
	"errors"
	"github.com/google/uuid"
)

type inMemoryRepository struct {
	users []*User
}

func NewInMemoryRepository() *inMemoryRepository {
	users := []*User{}
	return &inMemoryRepository{users}
}

func (r *inMemoryRepository) GetUsers(ctx context.Context) []*User {
	return r.users
}

func (r *inMemoryRepository) Create(ctx context.Context, name string) User {
	user := User{
		ID:    uuid.NewString(),
		Login: name,
	}
	r.users = append(r.users, &user)
	return user
}

func (r *inMemoryRepository) GetByID(ctx context.Context, ID string) (User, error) {
	for _, t := range r.users {
		if t.ID == ID {
			return *t, nil
		}
	}
	return User{}, errors.New("user is not found")
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
func remove(slice []*User, s int) []*User {
	return append(slice[:s], slice[s+1:]...)
}

func (r *inMemoryRepository) Update(ctx context.Context, ID string, Name string) (User, error) {
	for _, t := range r.users {
		if t.ID == ID {
			t.Login = Name
			return *t, nil
		}
	}
	return User{}, errors.New("user is not found")
}
