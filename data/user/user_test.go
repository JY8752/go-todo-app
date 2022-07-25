package user

import (
	"context"
	"os"
	"testing"
	"todo-app/ent"
	"todo-app/test"

	"github.com/stretchr/testify/assert"
)

var entClient *ent.Client
var userRepository *UserRepository

func TestMain(m *testing.M) {
	var terminateContainer func()
	terminateContainer, entClient = test.SetupMySQLContainer()
	defer terminateContainer()
	defer entClient.Close()

	userRepository = NewUserRepository(entClient)

	os.Exit(m.Run())
}

func TestFindUser(t *testing.T) {
	//given
	ctx := context.Background()
	created, _ := userRepository.Create(ctx, "user1", "test@test.com", 32)
	defer userRepository.DeleteAll(ctx)

	//when
	user, _ := userRepository.FindById(ctx, created.ID)

	//then
	assertUser(t, created, user)
}

func TestFindAll(t *testing.T) {
	//given
	ctx := context.Background()
	userRepository.Create(ctx, "user1", "test1@test.com", 20)
	userRepository.Create(ctx, "user2", "test2@test.com", 30)
	defer userRepository.DeleteAll(ctx)

	//when
	users, _ := userRepository.FindAll(ctx)

	//then
	assert.Equal(t, 2, len(users))
}

func assertUser(t *testing.T, expected *ent.User, actual *ent.User) {
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Email, actual.Email)
	assert.Equal(t, expected.Age, actual.Age)
}
