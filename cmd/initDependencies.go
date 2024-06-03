package cmd

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	controller "github.com/BrunoPolaski/go-crud/src/controller/user"
	"github.com/BrunoPolaski/go-crud/src/model/repository"
	"github.com/BrunoPolaski/go-crud/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	logger.Info("Setting up dependencies")
	repository := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repository)
	return controller.NewUserController(service)
}
