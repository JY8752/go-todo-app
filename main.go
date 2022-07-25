package main

import (
	"context"
	"fmt"
	"log"
	"todo-app/data"
	"todo-app/data/user"
)

func main() {
	ctx := context.Background()
	client := data.EntClient(ctx, "docker:docker@/todo")
	userRepository := user.NewUserRepository(client)
	defer client.Close()

	created, err := userRepository.Create(ctx, "user1", "text@text.com", 32)
	if err != nil {
		log.Fatalf("create error %v", err)
	}
	fmt.Printf("created: %v\n", created)
	get, _ := userRepository.FindById(ctx, created.ID)
	fmt.Printf("get: %v\n", get)
}
