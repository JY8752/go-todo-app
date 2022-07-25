package todo

import (
	"context"
	"time"
	"todo-app/ent"
	"todo-app/ent/todo"
)

type TodoRepository struct {
	client *ent.Client
}

func NewTodoRepository(c *ent.Client) *TodoRepository {
	return &TodoRepository{client: c}
}

func (repository TodoRepository) Create(ctx context.Context, userId int, title, detail string) (*ent.Todo, error) {
	return repository.client.Todo.Create().
		SetUserID(userId).
		SetTitle(title).
		SetDetail(detail).
		SetCreatedAt(time.Now().UTC()).
		SetUpdatedAt(time.Now().UTC()).
		Save(ctx)
}

func (repository TodoRepository) FindByUserId(ctx context.Context, userId int) ([]*ent.Todo, error) {
	return repository.client.Todo.Query().Where(todo.UserIDEQ(userId)).All(ctx)
}

func (repository TodoRepository) FindById(ctx context.Context, id int) (*ent.Todo, error) {
	return repository.client.Todo.Get(ctx, id)
}

func (repository TodoRepository) DeleteAll(ctx context.Context) (int, error) {
	return repository.client.Todo.Delete().Exec(ctx)
}
