package repository

import (
	"os"

	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/BrunoPolaski/go-crud/src/model/repository/entity/converter"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

func (ur *userRepository) UpdateUserRepository(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init UpdateUserRepository repository", zap.String("journey", "updateUser"))

	collection := ur.databaseConnection.Collection(os.Getenv(MONGO_USERS_DATABASE))

	entity := converter.ConvertDomainToEntity(userDomain)

	filter := bson.D{{Key: "_id", Value: entity.ID}}

	_, err := collection.UpdateOne()
}
