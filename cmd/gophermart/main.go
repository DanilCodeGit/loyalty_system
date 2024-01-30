package main

import (
	"context"
	"github.com/DanilCodeGit/loyalty_system/internal/conf"
	"github.com/DanilCodeGit/loyalty_system/internal/controller"
	"github.com/DanilCodeGit/loyalty_system/internal/database"
	"github.com/DanilCodeGit/loyalty_system/internal/repository"
	"github.com/DanilCodeGit/loyalty_system/internal/routes"
	"github.com/DanilCodeGit/loyalty_system/internal/service"
	"log"
	"net/http"
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
