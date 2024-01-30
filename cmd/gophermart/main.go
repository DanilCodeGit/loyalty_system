package main

import (
	"context"
	"log"
	"net/http"
	"v2/internal/conf"
	"v2/internal/controller"
	"v2/internal/database"
	"v2/internal/repository"
	"v2/internal/routes"
	"v2/internal/service"
)

func main() {
	ctx := context.Background()
	err := conf.InitConfig()
	if err != nil {
		log.Fatalf("init config: %s", err)
	}

	connDB, _ := database.NewDataBase(ctx, conf.DSN)
	userRepo := repository.NewUsers(connDB)
	userService := service.NewUsers(userRepo)
	controllerUser := controller.NewUsers(userService)
	c := routes.Controllers{
		UsersController: controllerUser,
	}

	r := routes.HandlersHTTP(c)
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
