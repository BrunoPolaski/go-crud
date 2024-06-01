package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/restErr"
	"github.com/BrunoPolaski/go-crud/src/model"
	"github.com/BrunoPolaski/go-crud/src/model/repository/entity/converter"
	"go.uber.org/zap"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *restErr.RestErr) {
	logger.Info("Init CreateUser repository", zap.String("journey", "createUser"))

	collectionName := os.Getenv("MONGO_USERS_DATABASE")
	collection := ur.databaseConnection.Collection(collectionName)

	entity := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), entity)
	if err != nil {
		return nil, restErr.NewInternalServerError(fmt.Sprintf("Error creating user %v", err))
	}

	userDomain.SetId(result.InsertedID.(string))

	return userDomain, nil
}