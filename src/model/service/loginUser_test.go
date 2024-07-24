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
	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)
	user := mocks.NewMockUserDomainInterface(ctrl)

	t.Run("shall_return_success_when_credentials_are_valid", func(t *testing.T) {
		user.EXPECT().GetEmail().Return("email")
		user.EXPECT().GetPassword().Return("password")

		repository.EXPECT().FindUserByEmailRepository(user.GetEmail()).Return(
			user,
			nil,
		)

		user.EXPECT().ComparePassword(user.GetPassword()).Return(
			nil,
		)

		user.EXPECT().GenerateToken().Return(
			"token",
			nil,
		)

		user.EXPECT().GetEmail().Return("email")
		user.EXPECT().GetPassword().Return("password")

		response, token, err := service.LoginUserService(user)

		assert.Nil(t, err)
		assert.NotNil(t, response)
		assert.NotNil(t, token)
		assert.Equal(t, "token", token)
		assert.Equal(t, user, response)
	})

	t.Run("shall_return_error_when_user_not_found", func(t *testing.T) {
		user.EXPECT().GetEmail().Return("email")
		repository.EXPECT().FindUserByEmailRepository(user.GetEmail()).Return(
			nil,
			rest_err.NewNotFoundError("User not found"),
		)

		user.EXPECT().GetEmail().Return("email")

		user, token, err := service.LoginUserService(user)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.Equal(t, token, "")
	})

	t.Run("shall_return_error_when_user_not_found", func(t *testing.T) {
		user.EXPECT().GetEmail().Return("email")
		repository.EXPECT().FindUserByEmailRepository(user.GetEmail()).Return(
			nil,
			rest_err.NewNotFoundError("User not found"),
		)

		user.EXPECT().GetEmail().Return("email")

		user, token, err := service.LoginUserService(user)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.Equal(t, token, "")
	})
}
