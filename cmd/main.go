package main

import (
	"log"

	todo "github.com/BudjakovDmitry/go_todo_app"
	"github.com/BudjakovDmitry/go_todo_app/pkg/handler"
	"github.com/BudjakovDmitry/go_todo_app/pkg/repository"
	"github.com/BudjakovDmitry/go_todo_app/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s\n", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewJandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occupied while running http server: %s\n", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
