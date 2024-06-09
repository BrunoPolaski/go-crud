package repository

import (
	"context"

	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUserRepository(id string) *rest_err.RestErr {
	logger.Info("Init DeleteUserRepository repository", zap.String("journey", "DeleteUserRepository"))

	collection := ur.databaseConnection.Collection(MONGO_USERS_DATABASE)

	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		logger.Error(
			"Error trying to delete user", err,
			zap.String("journey", "DeleteUserRepository"),
			zap.String("id", id),
		)
		return rest_err.NewInternalServerError("Error trying to delete user")
	}

	return nil
}
