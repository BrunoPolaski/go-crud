package main

import (
	"context"
	"log"

	"github.com/BrunoPolaski/go-crud/src/configuration"
	"github.com/BrunoPolaski/go-crud/src/configuration/database/mongodb"
	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting the application")

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	database, err := mongodb.NewMongoConnection(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	userController := configuration.InitDependencies(database)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	routes.InitRoutes(userController, &router.RouterGroup)

	logger.Info("All set to Go!")
	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
