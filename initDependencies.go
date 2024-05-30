package main

import (
	"github.com/BrunoPolaski/go-crud/src/controller"
	"github.com/BrunoPolaski/go-crud/src/model/repository"
	"github.com/BrunoPolaski/go-crud/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	repository := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repository)
	return controller.NewUserController(service)
}
