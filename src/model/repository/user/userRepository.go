package repository

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGO_USERS_DATABASE = "MONGO_USERS_DATABASE"
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
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUserRepository(userDomain model.UserDomainInterface, id string) *rest_err.RestErr
	DeleteUserRepository(id string) *rest_err.RestErr
}
