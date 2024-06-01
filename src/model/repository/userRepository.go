package repository

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/restErr"
	"github.com/BrunoPolaski/go-crud/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(databaseConnection *mongo.Database) UserRepository {
	return &userRepository{
		databaseConnection: databaseConnection,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *restErr.RestErr)
}
