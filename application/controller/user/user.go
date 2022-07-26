package user

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"

	"todo-app/application/controller"
	data "todo-app/data/user"
	domain "todo-app/domain/service/user"
	"todo-app/ent"
)

func UserRoute(ctx context.Context, entClient *ent.Client) {
	service := domain.UserService{UserRepository: data.NewUserRepository(entClient)}

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) { userHandle(ctx, service, w, r) })
	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) { userHandle(ctx, service, w, r) })
}

func userHandle(ctx context.Context, service domain.UserService, w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		id, err := controller.GetPathParameter(r.URL.Path, "/user")
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}

		user, err := service.GetUser(ctx, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		writeUserResponse(user, w)

	case http.MethodPost:
		var params struct {
			Name  string `json:"name"`
			Email string `json:"email"`
			Age   int    `json:"age"`
		}

		controller.JsonDecode(r, w, &params)

		user, err := service.Create(ctx, params.Name, params.Email, params.Age)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		writeUserResponse(user, w)
	default:
		log.Printf("unkown http method %v", r.Method)
	}
}

func writeUserResponse(user *ent.User, w http.ResponseWriter) string {
	var buf bytes.Buffer
	if err := controller.JsonEncode(&buf, user); err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, buf.String())

	return buf.String()
}
