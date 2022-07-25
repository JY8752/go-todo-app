package todo

import (
	"context"
	"testing"
	"todo-app/ent"
	mock_todo "todo-app/mock/domain/service/todo"
	"todo-app/test"

	"github.com/golang/mock/gomock"
)

func TestCreate(t *testing.T) {
	//given
	ctx := context.Background()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	expected := &ent.Todo{ID: 1, UserID: 1, Title: "title", Detail: "detail"}
	mockRep := mock_todo.NewMockTodoReposityory(mockCtrl)
	mockRep.EXPECT().Create(ctx, 1, "title", "detail").Return(expected, nil)

	todoService := &TodoService{todoRepository: mockRep}

	//when
	result, _ := todoService.Create(ctx, 1, "title", "detail")

	//then
	test.AssertTodo(t, expected, result, false)
}
