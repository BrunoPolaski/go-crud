package service

import (
	"testing"

	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/BrunoPolaski/go-crud/src/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_FindUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(mockRepository)

	t.Run("shall_return_success_when_user_exists", func(t *testing.T) {
		userDomain := mocks.UserMock

		mockRepository.EXPECT().FindUserByEmailRepository(userDomain.GetEmail()).Return(
			userDomain,
			nil,
		)

		response, err := service.FindUserByEmailService(userDomain.GetEmail())

		assert.NotNil(t, response)
		assert.Nil(t, err)
		assert.Equal(t, response, userDomain)
	})

	t.Run("shall_return_error_when_user_not_found", func(t *testing.T) {

		mockRepository.EXPECT().FindUserByEmailRepository("test").Return(
			nil,
			rest_err.NewNotFoundError("user not found"),
		)

		response, err := service.FindUserByEmailService("test")

		assert.NotNil(t, err)
		assert.Nil(t, response)
		assert.Equal(t, err.Message, "user not found")
	})
}

func TestUserService_FindUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(mockRepository)

	t.Run("shall_return_success_when_user_exists", func(t *testing.T) {
		userDomain := mocks.UserMock

		mockRepository.EXPECT().FindUserByIDRepository("1").Return(
			userDomain,
			nil,
		)

		response, err := service.FindUserByIDService("1")

		assert.NotNil(t, response)
		assert.Nil(t, err)
		assert.Equal(t, response, userDomain)
	})

	t.Run("shall_return_error_when_user_not_found", func(t *testing.T) {

		mockRepository.EXPECT().FindUserByIDRepository("test").Return(
			nil,
			rest_err.NewNotFoundError("user not found"),
		)

		response, err := service.FindUserByIDService("test")

		assert.NotNil(t, err)
		assert.Nil(t, response)
		assert.Equal(t, err.Message, "user not found")
	})
}
