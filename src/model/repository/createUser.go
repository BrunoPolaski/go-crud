package repository

import (
	"context"
	"fmt"

	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/BrunoPolaski/go-crud/src/model/repository/entity/converter"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateUser repository", zap.String("journey", "createUser"))

	collection := ur.databaseConnection.Collection(collectionName)

	entity := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), entity)
	if err != nil {
		return nil, rest_err.NewInternalServerError(fmt.Sprintf("Error creating user %v", err))
	}

	entity.ID = result.InsertedID.(primitive.ObjectID)

	return converter.ConvertEntityToDomain(entity), nil
}