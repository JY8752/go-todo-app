package todo

import (
	"context"
	"os"
	"testing"
	"todo-app/data/user"
	"todo-app/ent"
	"todo-app/test"

	"github.com/stretchr/testify/assert"
)

var entClient *ent.Client
var userRepository *user.UserRepository
var todoRepository *TodoRepository

func TestMain(m *testing.M) {
	var terminateContainer func()
	terminateContainer, entClient = test.SetupMySQLContainer()
	defer terminateContainer()
	defer entClient.Close()

	todoRepository = NewTodoRepository(entClient)
	userRepository = user.NewUserRepository(entClient)

	os.Exit(m.Run())
}

func TestFindTodo(t *testing.T) {
	//given
	ctx := context.Background()
	user, _ := userRepository.Create(ctx, "user1", "test@test.com", 32)
	created, _ := todoRepository.Create(ctx, user.ID, "title1", "detail1")
	defer userRepository.DeleteAll(ctx)
	defer todoRepository.DeleteAll(ctx)

	//when
	todos, _ := todoRepository.FindByUserId(ctx, user.ID)

	//then
	assert.Equal(t, 1, len(todos))
	test.AssertTodo(t, created, todos[0], true)
}
