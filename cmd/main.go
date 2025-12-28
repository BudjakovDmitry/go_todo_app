package main

import (
	"log"

	todo "github.com/BudjakovDmitry/go_todo_app"
	"github.com/BudjakovDmitry/go_todo_app/pkg/handler"
	"github.com/BudjakovDmitry/go_todo_app/pkg/repository"
	"github.com/BudjakovDmitry/go_todo_app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewJandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal("error occupied while running http server: %s", err.Error())
	}
}
