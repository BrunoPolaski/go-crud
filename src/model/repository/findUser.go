package repository

import (
	"context"

	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/BrunoPolaski/go-crud/src/model/repository/entity"
	"github.com/BrunoPolaski/go-crud/src/model/repository/entity/converter"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
	"go.mongodb.org/mongo-driver/bson"
)

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Initiating FindUser repository")

	collection := ur.databaseConnection.Collection(collectionName)

	entity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}

	err := collection.FindOne(context.Background(), filter).Decode(entity)
	if err != nil {
		return nil, rest_err.NewInternalServerError("Error finding user")
	}

	return converter.ConvertEntityToDomain(entity), nil
}
