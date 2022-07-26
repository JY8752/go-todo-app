package user

import (
	"context"
	// "todo-app/data/user"
	"todo-app/ent"
)

type UserRepository interface {
	// NewUserRepository(*ent.Client) *user.UserRepository
	FindById(context.Context, int) (*ent.User, error)
	FindAll(context.Context) ([]*ent.User, error)
	Create(context.Context, string, string, int) (*ent.User, error)
}

type UserService struct {
	UserRepository UserRepository
}

func (service UserService) Create(ctx context.Context, name, email string, age int) (*ent.User, error) {
	return service.UserRepository.Create(ctx, name, email, age)
}
