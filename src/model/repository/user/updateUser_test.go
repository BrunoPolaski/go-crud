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

func TestUserRepository_UpdateUser(t *testing.T) {
	databaseName := "user_database_test"
	collectionName := "user_collection_test"

	err := os.Setenv("MONGODB_USER_DB", collectionName)
	if err != nil {
		t.FailNow()
		return
	}
	defer os.Clearenv()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.ClearCollections()

	mt.Run("shall_return_success_when_sending_a_valid_domain", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})
		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		domain := model.NewUserDomain(
			"test@test.com", "test", "test", 90)
		err := repo.UpdateUserRepository(domain, primitive.NewObjectID().Hex())

		assert.Nil(t, err)
	})

	mt.Run("shall__return_error_when_database_returns_error", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})
		databaseMock := mt.Client.Database(databaseName)

		repo := NewUserRepository(databaseMock)
		domain := model.NewUserDomain(
			"test@test.com", "test", "test", 90)
		err := repo.UpdateUserRepository(domain, primitive.NewObjectID().Hex())

		assert.NotNil(t, err)
	})
}
