package repository

import (
	"context"
	"os"

	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/BrunoPolaski/go-crud/src/model/repository/entity/converter"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateUserRepository(userDomain model.UserDomainInterface, id string) *rest_err.RestErr {
	logger.Info("Init UpdateUserRepository repository", zap.String("journey", "updateUser"))

	collection := ur.databaseConnection.Collection(os.Getenv(MONGO_USERS_DATABASE))

	entity := converter.ConvertDomainToEntity(userDomain)
	primitiveId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error("Error converting id to primitive.ObjectID", err,
			zap.String("journey", "updateUser"),
		)
		return rest_err.NewBadRequestError("User ID is not valid")
	}

	filter := bson.D{{Key: "_id", Value: primitiveId}}

	_, err = collection.UpdateOne(context.Background(), filter, bson.D{{Key: "$set", Value: entity}})
	if err != nil {
		logger.Error("Error updating user", err,
			zap.String("journey", "updateUser"),
		)
		return rest_err.NewInternalServerError("Error updating user")
	} else {
		return nil
	}
}
