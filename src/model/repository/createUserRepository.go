package repository

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/BrunoPolaski/go-crud/src/model"
	"go.uber.org/zap"
)

func (ur *userRepository) CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateUser repository", zap.String("journey", "createUser"))

	// collectionName := os.Getenv("MONGO_USERS_DATABASE")
	// collection := ur.databaseConnection.Collection(collectionName)

	//collection.InsertOne(context.Background(), userDomain)

	return nil, nil
}
