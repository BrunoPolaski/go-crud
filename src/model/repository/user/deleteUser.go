package repository

import (
	"context"
	"os"

	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUserRepository(id string) *rest_err.RestErr {
	logger.Info("Init DeleteUserRepository repository", zap.String("journey", "DeleteUserRepository"))

	collection := ur.databaseConnection.Collection(os.Getenv(MONGO_USERS_COLLECTION))

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}

	if _, err := collection.DeleteOne(context.Background(), filter); err != nil {
		logger.Error(
			"Error trying to delete user", err,
			zap.String("journey", "DeleteUserRepository"),
			zap.String("id", id),
		)
		return rest_err.NewInternalServerError("Error trying to delete user")
	}

	logger.Info("Successfully executed DeleteUserRepository repository",
		zap.String("journey", "DeleteUserRepository"),
		zap.String("id", id),
	)
	return nil
}
