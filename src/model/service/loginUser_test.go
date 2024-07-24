package service

import (
	"testing"

	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/BrunoPolaski/go-crud/src/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_LoginUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockRepository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(mockRepository)
	mockUser := mocks.NewMockUserDomainInterface(ctrl)

	t.Run("shall_return_success_when_credentials_are_valid", func(t *testing.T) {
		mockUser.EXPECT().GetEmail().Return("email")
		mockUser.EXPECT().GetPassword().Return("password")

		mockRepository.EXPECT().FindUserByEmailRepository(mockUser.GetEmail()).Return(
			mockUser,
			nil,
		)

		mockUser.EXPECT().ComparePassword(mockUser.GetPassword()).Return(
			nil,
		)

		mockUser.EXPECT().GenerateToken().Return(
			"token",
			nil,
		)

		mockUser.EXPECT().GetEmail().Return("email")
		mockUser.EXPECT().GetPassword().Return("password")

		response, token, err := service.LoginUserService(mockUser)

		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, token)
		assert.Equal(t, "token", token)
		assert.Equal(t, mockUser, response)
	})

	t.Run("shall_return_error_when_user_not_found", func(t *testing.T) {
		mockUser.EXPECT().GetEmail().Return("email")
		mockRepository.EXPECT().FindUserByEmailRepository(mockUser.GetEmail()).Return(
			nil,
			rest_err.NewNotFoundError("mockUser not found"),
		)

		mockUser.EXPECT().GetEmail().Return("email")

		mockUser, token, err := service.LoginUserService(mockUser)

		assert.Nil(t, mockUser)
		assert.NotNil(t, err)
		assert.Equal(t, token, "")
	})

	t.Run("shall_return_error_when_passwords_mismatch", func(t *testing.T) {
		mockUser.EXPECT().GetEmail().AnyTimes().Return("email")
		mockUser.EXPECT().GetPassword().AnyTimes().Return("password")
		mockRepository.EXPECT().FindUserByEmailRepository("email").Return(
			mockUser,
			nil,
		)

		mockUser.EXPECT().ComparePassword("password").Return(
			rest_err.NewUnauthorizedError("Invalid credentials"),
		)

		mockUser, token, err := service.LoginUserService(mockUser)

		assert.Nil(t, mockUser)
		assert.NotNil(t, err)
		assert.Equal(t, token, "")
		assert.Equal(t, err.Error(), "Invalid credentials")
	})

	t.Run("shall_return_error_when_not_generating_token", func(t *testing.T) {
		mockUser.EXPECT().GetEmail().AnyTimes().Return("email")
		mockUser.EXPECT().GetPassword().AnyTimes().Return("password")
		mockRepository.EXPECT().FindUserByEmailRepository("email").Return(
			mockUser,
			nil,
		)

		mockUser.EXPECT().ComparePassword("password").Return(
			nil,
		)

		mockUser.EXPECT().GenerateToken().Return(
			"",
			rest_err.NewInternalServerError("Error trying to create jwt"),
		)

		mockUser, token, err := service.LoginUserService(mockUser)

		assert.Nil(t, mockUser)
		assert.NotNil(t, err)
		assert.Equal(t, token, "")
		assert.Equal(t, err.Error(), "Error trying to create jwt")
	})
}
