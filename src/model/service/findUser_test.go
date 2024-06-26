package service

import (
	"testing"

	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
	"github.com/BrunoPolaski/go-crud/src/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_FindUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("shall_return_success_when_user_exists", func(t *testing.T) {
		userDomain := model.NewUserDomain(
			"test@test.com",
			"test",
			"bruno",
			19,
		)

		repository.EXPECT().FindUserByEmailRepository(userDomain.GetEmail()).Return(
			userDomain,
			nil,
		)

		response, err := service.FindUserByEmailService(userDomain.GetEmail())

		assert.NotNil(t, response)
		assert.Nil(t, err)
		assert.Equal(t, response, userDomain)
	})

	t.Run("shall_return_error", func(t *testing.T) {

		repository.EXPECT().FindUserByEmailRepository("test").Return(
			nil,
			rest_err.NewNotFoundError("user not found"),
		)

		response, err := service.FindUserByEmailService("test")

		assert.NotNil(t, err)
		assert.Nil(t, response)
		assert.Equal(t, err.Message, "user not found")
	})
}
