package user

import (
	"context"
	"testing"
	"todo-app/ent"
	mock_user "todo-app/mock/domain/service/user"
	"todo-app/test"

	"github.com/golang/mock/gomock"
)

func TestCreate(t *testing.T) {
	//given
	ctx := context.Background()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	expected := &ent.User{ID: 1, Name: "user1", Email: "test@test.com", Age: 32}
	mockRep := mock_user.NewMockUserRepository(mockCtrl)
	mockRep.EXPECT().Create(ctx, "user1", "test@text,com", 32).Return(expected, nil)

	userService := &UserService{UserRepository: mockRep}

	//when
	rerult, _ := userService.Create(ctx, "user1", "test@text,com", 32)

	//then
	test.AssertUser(t, expected, rerult)
}
