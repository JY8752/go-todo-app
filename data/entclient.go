package data

import (
	"context"
	"fmt"
	"log"
	"todo-app/ent"
	"todo-app/ent/migrate"

	_ "github.com/go-sql-driver/mysql"
)

func EntClient(ctx context.Context, connectionString string) *ent.Client {
	// client, err := ent.Open("mysql", "docker:docker@/todo
	client, err := ent.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}

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

func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if v := recover(); v != nil {
			tx.Rollback()
			panic(v)
		}
	}()
	if err := fn(tx); err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("rolling back transaction: %w", rerr)
		}
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
