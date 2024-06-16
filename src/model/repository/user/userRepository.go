package repository

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGO_USERS_COLLECTION = "MONGO_USERS_COLLECTION"
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
	CreateUserRepository(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailRepository(email string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByIDRepository(id string) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUserRepository(userDomain model.UserDomainInterface, id string) *rest_err.RestErr
	DeleteUserRepository(id string) *rest_err.RestErr
}
