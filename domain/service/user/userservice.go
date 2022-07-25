package user

import (
	"context"
	"todo-app/data/user"
	"todo-app/ent"
)

type UserService struct {
	userRepository *user.UserRepository
}

func (service UserService) Create(ctx context.Context, name, email string, age int) (*ent.User, error) {
	return service.userRepository.Create(ctx, name, email, age)
}
