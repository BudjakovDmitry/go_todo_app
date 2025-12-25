package main

import (
	todo "github.com/BudjakovDmitry/go_todo_app"
	"github.com/BudjakovDmitry/go_todo_app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal("error occupied while running http server: %s", err.Error())
	}
}
