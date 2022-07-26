package todo

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"todo-app/application/controller"
	"todo-app/ent"

	data "todo-app/data/todo"
	domain "todo-app/domain/service/todo"
)

func TodoRoute(ctx context.Context, entClient *ent.Client) {
	service := domain.TodoService{TodoRepository: data.NewTodoRepository(entClient)}

	http.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var params struct {
				UserId int    `json:"user_id"`
				Title  string `json:"title"`
				Detail string `json:"detail"`
			}

			controller.JsonDecode(r, w, &params)

			todo, err := service.Create(ctx, params.UserId, params.Title, params.Detail)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			var buf bytes.Buffer
			controller.JsonEncode(&buf, todo)

			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, buf.String())

		default:
			log.Printf("unkown http method %v", r.Method)
		}
	})
}
