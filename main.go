package main

import (
	"context"
	"log"
	"net/http"

	"todo-app/application"
	"todo-app/data"
)

func main() {
	ctx := context.Background()
	client := data.EntClient(ctx, "root:root@tcp(localhost:3306)/todo-app?parseTime=true")

	application.Route(ctx, client)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
