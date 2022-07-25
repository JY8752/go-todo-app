package todo

import (
	"context"
	"todo-app/data/todo"
	"todo-app/ent"
)

type TodoService struct {
	todoRepository *todo.TodoRepository
}

func (service TodoService) Create(ctx context.Context, userId int, title, detail string) (*ent.Todo, error) {
	return service.todoRepository.Create(ctx, userId, title, detail)
}
