package repository

import (
	"os"
	"testing"

	model "github.com/BrunoPolaski/go-crud/src/model/user"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_CreateUser(t *testing.T) {
	databaseName := "user_database_test"
	collection := "user_collection_test"

	os.Setenv("MONGO_USERS_COLLECTION", collection)

	mtestDB := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mtestDB.Run("when_sending_valid_domain_shall_return_success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: "1"},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})
		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.CreateUserRepository(model.NewUserDomain(
			"test@test.com",
			"test",
			"test",
			90,
		))

		_, idErr := primitive.ObjectIDFromHex(userDomain.GetID())

		assert.Nil(t, err)
		assert.Nil(t, idErr)
		assert.EqualValues(t, userDomain.GetEmail(), "test@test.com")
	})

	defer mtestDB.Cleanup(mtestDB.ClearCollections)
	defer os.Clearenv()
}
