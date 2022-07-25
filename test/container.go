package test

import (
	"context"
	"fmt"
	"log"
	"todo-app/data"
	"todo-app/ent"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUsername string = "root"
	dbPassword string = "password"
	dbName     string = "test"
)

func SetupMySQLContainer() (func(), *ent.Client) {
	log.Println("setup MySQL Container")
	ctx := context.Background()
	mysqlPort, _ := nat.NewPort("tcp", "3306")

	req := testcontainers.ContainerRequest{
		Image:        "mysql:latest",
		ExposedPorts: []string{"3306/tcp"},
		Env: map[string]string{
			"MYSQL_ROOT_PASSWORD": dbPassword,
			"MYSQL_DATABASE":      dbName,
		},
		WaitingFor: wait.ForListeningPort(mysqlPort),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	closeContainer := func() {
		log.Println("terminating container")
		err := container.Terminate(ctx)
		if err != nil {
			panic(fmt.Sprintf("%v", err))
		}
	}

	host, _ := container.Host(ctx)
	p, _ := container.MappedPort(ctx, "3306/tcp")
	port := p.Int()

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		dbUsername, dbPassword, host, port, dbName)

	log.Println(connectionString)

	entClient := data.EntClient(ctx, connectionString)

	return closeContainer, entClient
}
