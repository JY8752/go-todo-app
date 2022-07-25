package user

import (
	"context"
	"todo-app/ent"

	_ "github.com/go-sql-driver/mysql"
)

type UserRepository struct {
	client *ent.Client
}

func NewUserRepository(c *ent.Client) *UserRepository {
	return &UserRepository{client: c}
}

func (repository UserRepository) FindById(ctx context.Context, id int) (*ent.User, error) {
	return repository.client.User.Get(ctx, id)
}

func (repository UserRepository) FindAll(ctx context.Context) ([]*ent.User, error) {
	return repository.client.User.Query().All(ctx)
}

func (repository UserRepository) Create(ctx context.Context, name, email string, age int) (*ent.User, error) {
	return repository.client.User.Create().
		SetName(name).
		SetEmail(email).
		SetAge(age).
		Save(ctx)
}

func (repository UserRepository) DeleteAll(ctx context.Context) (int, error) {
	return repository.client.User.Delete().Exec(ctx)
}
