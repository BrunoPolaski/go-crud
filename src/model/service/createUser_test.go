package service

import (
	"testing"

	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/BrunoPolaski/go-crud/src/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserService_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(mockRepository)

	t.Run("shall_return_user_when_success", func(t *testing.T) {
		userMock := mocks.UserMock

		mockRepository.EXPECT().FindUserByEmailRepository(userMock.GetEmail()).Return(nil, nil)

		mockRepository.EXPECT().CreateUserRepository(userMock).Return(userMock, nil)

		response, err := service.CreateUserService(userMock)

		assert.NotNil(t, response, "Expected non-nil response")
		assert.Nil(t, err, "Expected nil error")
	})

	t.Run("shall_return_error_when_user_exists", func(t *testing.T) {
		userMock := mocks.UserMock

		userMock.SetID(primitive.NewObjectID().Hex())

		mockRepository.EXPECT().FindUserByEmailRepository(userMock.GetEmail()).Return(
			userMock,
			nil,
		)

		response, err := service.CreateUserService(userMock)

		assert.NotNil(t, err)
		assert.Nil(t, response)
		assert.Equal(t, err.Message, "User already exists")
	})

	t.Run("shall_return_error_when_internal_error", func(t *testing.T) {
		userMock := mocks.UserMock

		userMock.SetID(primitive.NewObjectID().Hex())

		mockRepository.EXPECT().FindUserByEmailRepository(userMock.GetEmail()).Return(
			nil,
			nil,
		)

		mockRepository.EXPECT().CreateUserRepository(userMock).Return(
			nil,
			rest_err.NewInternalServerError("Error creating user"),
		)

		response, err := service.CreateUserService(userMock)

		assert.NotNil(t, err)
		assert.Nil(t, response)
		assert.Equal(t, err.Message, "Error creating user")
	})
}
