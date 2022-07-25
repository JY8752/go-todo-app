package todo

import (
	"context"
	"os"
	"testing"
	"time"
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
	assertTodo(t, created, todos[0])
}

func assertTodo(t *testing.T, expected, actual *ent.Todo) {
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.UserID, actual.UserID)
	assert.Equal(t, expected.Title, actual.Title)
	assert.Equal(t, expected.Detail, actual.Detail)
	assertTime(t, expected.CreatedAt, actual.CreatedAt)
	assertTime(t, expected.UpdatedAt, actual.UpdatedAt)
	assertTime(t, expected.CompletedAt, actual.CompletedAt)
}

func assertTime(t *testing.T, expected, actual time.Time) {
	type AssertTime struct {
		year   int
		mounth int
		day    int
		hour   int
		minute int
	}
	expectedTime := AssertTime{
		year:   expected.Year(),
		mounth: int(expected.Month()),
		day:    expected.Day(),
		hour:   expected.Hour(),
		minute: expected.Minute(),
	}
	actualTime := AssertTime{
		year:   actual.Year(),
		mounth: int(actual.Month()),
		day:    actual.Day(),
		hour:   actual.Hour(),
		minute: actual.Minute(),
	}
	assert.Equal(t, expectedTime, actualTime)
}
