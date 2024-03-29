package todo

import (
	"context"
	"todo-app/ent"
)

type TodoReposityory interface {
	Create(context.Context, int, string, string) (*ent.Todo, error)
	FindByUserId(context.Context, int) ([]*ent.Todo, error)
	FindById(context.Context, int) (*ent.Todo, error)
}

type TodoService struct {
	TodoRepository TodoReposityory
}

func (service TodoService) Create(ctx context.Context, userId int, title, detail string) (*ent.Todo, error) {
	return service.TodoRepository.Create(ctx, userId, title, detail)
}
