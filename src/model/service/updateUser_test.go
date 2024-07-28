package service

import (
	"testing"

	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/BrunoPolaski/go-crud/src/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserService_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(mockRepository)

	t.Run("shall_return_success_when_payload_is_correct", func(t *testing.T) {
		userDomain := mocks.UserMock

		id := primitive.NewObjectID().Hex()

		mockRepository.EXPECT().UpdateUserRepository(userDomain, id).Return(
			nil,
		)

		err := service.UpdateUserService(userDomain, id)

		assert.Nil(t, err)
	})

	t.Run("shall_return_error_when_internal_error", func(t *testing.T) {
		userDomain := mocks.UserMock

		id := primitive.NewObjectID().Hex()

		mockRepository.EXPECT().UpdateUserRepository(userDomain, id).Return(
			rest_err.NewInternalServerError("Error updating user"),
		)

		err := service.UpdateUserService(userDomain, id)

		assert.NotNil(t, err)
		assert.Equal(t, rest_err.NewInternalServerError("Error updating user"), err)
	})
}
