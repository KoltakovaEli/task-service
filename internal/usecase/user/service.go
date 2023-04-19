package user

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
type userService struct {
	taskRepo task.Repo
	userRepo user.Repo
}

func NewService(ur user.Repo) UseCase {
	return &userService{
		userRepo: ur,
	}
}

func (u *userService) GetUsers(ctx context.Context) ([]*entity.User, error) {
	return u.userRepo.GetUsers(ctx)
}

func (u *userService) Create(ctx context.Context, login, email string) (*entity.User, error) {
	return u.userRepo.Create(ctx, login, email)
}

func (u *userService) GetByID(ctx context.Context, ID string) (*entity.User, error) {
	return u.userRepo.GetByID(ctx, ID)
}

func (u *userService) Delete(ctx context.Context, ID string) error {
	return u.userRepo.Delete(ctx, ID)
}

func (u *userService) Update(ctx context.Context, ID, login, email string) (*entity.User, error) {
	return u.userRepo.Update(ctx, ID, login, email)
}
