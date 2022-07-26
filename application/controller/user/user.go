package user

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	data "todo-app/data/user"
	domain "todo-app/domain/service/user"
	"todo-app/ent"
)

func UserRoute(ctx context.Context, entClient *ent.Client) {
	service := domain.UserService{UserRepository: data.NewUserRepository(entClient)}

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
		case http.MethodPost:
			var params struct {
				Name  string `json:"name"`
				Email string `json:"email"`
				Age   int    `json:"age"`
			}

			dec := json.NewDecoder(r.Body)
			if err := dec.Decode(&params); err != nil {
				http.Error(w, err.Error(), 400)
			}

			user, err := service.Create(ctx, params.Name, params.Email, params.Age)
			if err != nil {
				http.Error(w, err.Error(), 500)
			}

			var buf bytes.Buffer
			enc := json.NewEncoder(&buf)
			if err = enc.Encode(user); err != nil {
				http.Error(w, err.Error(), 500)
			}

			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, buf.String())
		default:
			log.Printf("unkown http method %v", r.Method)
		}
	})
}
