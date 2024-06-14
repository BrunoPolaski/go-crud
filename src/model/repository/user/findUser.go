package repository

import (
	"context"
	"os"

	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/BrunoPolaski/go-crud/src/model/repository/entity"
	"github.com/BrunoPolaski/go-crud/src/model/repository/entity/converter"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ur *userRepository) FindUserByIDRepository(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Initiating FindUserByID repository")

	collection := ur.databaseConnection.Collection(os.Getenv(MONGO_USERS_COLLECTION))

	entity := &entity.UserEntity{}

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}

	err := collection.FindOne(context.Background(), filter).Decode(entity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError("User not found")
		}
		return nil, rest_err.NewInternalServerError("Error finding user by ID")
	}

	return converter.ConvertEntityToDomain(entity), nil
}

func (ur *userRepository) FindUserByEmailRepository(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Initiating FindUserByEmail repository")

	collection := ur.databaseConnection.Collection(os.Getenv(MONGO_USERS_COLLECTION))

	entity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}

	err := collection.FindOne(context.Background(), filter).Decode(entity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError("User not found")
		}
		return nil, rest_err.NewInternalServerError("Error finding user by email")
	}

	return converter.ConvertEntityToDomain(entity), nil
}

func (ur *userRepository) FindUserByPasswordRepository(password string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Initiating FindUserByPassword repository")

	collection := ur.databaseConnection.Collection(os.Getenv(MONGO_USERS_COLLECTION))

	entity := &entity.UserEntity{}

	filter := bson.D{{Key: "password", Value: password}}

	err := collection.FindOne(context.Background(), filter).Decode(entity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError("User not found")
		}
		return nil, rest_err.NewInternalServerError("Error finding user by password")
	}

	return converter.ConvertEntityToDomain(entity), nil
}
