package application

import (
	"context"
	"fmt"
	"todo-app/application/controller/user"
	"todo-app/ent"
)

func Route(ctx context.Context, ent *ent.Client) {
	fmt.Println("initialize router.")

	//user route
	user.UserRoute(ctx, ent)
}
