package data

import (
	"context"
	"log"
	"todo-app/ent"
	"todo-app/ent/migrate"

	_ "github.com/go-sql-driver/mysql"
)

func EntClient() *ent.Client {
	client, err := ent.Open("mysql", "docker:docker@/todo-app")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()

	ctx := context.Background()
	//マイグレーションの実行
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
