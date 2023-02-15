package user

import "context"

type Repo interface {
	GetUsers(ctx context.Context) ([]*User, error)
	Create(ctx context.Context, login string, email string) (User, error)
	GetByID(ctx context.Context, ID string) (User, error)
	Delete(ctx context.Context, ID string) error
	Update(ctx context.Context, ID string, login string, email string) (User, error)
}

type User struct {
	ID    string `json:"id"`
	Login string `json:"login"`
	Email string `json:"email"`
}
