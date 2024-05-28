package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/controller"
	"github.com/BrunoPolaski/go-crud/src/controller/routes"
	"github.com/BrunoPolaski/go-crud/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func throwError() (int, error) {
	return 0, errors.New("error")
}

func main() {
	logger.Info("Starting the application")

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	//Init dependencies
	service := service.NewUserDomainService()
	controller := controller.NewUserController(service)

	router := gin.Default()

	routes.InitRoutes(controller, &router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

	if _, err := throwError(); err != nil {
		fmt.Println(err)
	}
}
