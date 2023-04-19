package user

import (
	"context"
	"task-service/internal/entity"
)

type Repo interface {
	GetUsers(ctx context.Context) ([]*entity.User, error)
	Create(ctx context.Context, login string, email string) (*entity.User, error)
	GetByID(ctx context.Context, ID string) (*entity.User, error)
	Delete(ctx context.Context, ID string) error
	Update(ctx context.Context, ID string, login string, email string) (*entity.User, error)
}
