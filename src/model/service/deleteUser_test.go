package service

import (
	"testing"

	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/BrunoPolaski/go-crud/src/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(mockRepository)

	t.Run("shall_return_success_when_user_deleted", func(t *testing.T) {
		mockRepository.EXPECT().DeleteUserRepository("1").Return(
			nil,
		)

		err := service.DeleteUserService("1")

		assert.Nil(t, err)
	})

	t.Run("shall_return_error_when_internal_error", func(t *testing.T) {
		mockRepository.EXPECT().DeleteUserRepository("1").Return(
			rest_err.NewInternalServerError("Error trying to delete user"),
		)

		err := service.DeleteUserService("1")

		assert.NotNil(t, err)
		assert.Equal(t, err.Message, "Error trying to delete user")
	})
}
