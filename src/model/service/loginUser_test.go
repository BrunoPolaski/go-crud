package service

// import (
// 	"testing"

// 	"github.com/BrunoPolaski/go-crud/src/tests/mocks"
// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/assert"
// )

// func TestUserService_LoginUser(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	repository := mocks.NewMockUserRepository(ctrl)
// 	service := NewUserDomainService(repository)

// 	t.Run("shall_return_success_when_credentials_are_valid", func(t *testing.T) {
// 		userDomain := mocks.UserMock

// 		repository.EXPECT().FindUserByEmailRepository(userDomain.GetEmail()).Return(
// 			userDomain,
// 			nil,
// 		)

// 		response, token, err := service.LoginUserService(userDomain)

// 		assert.Nil(t, err)
// 		assert.NotNil(t, response)
// 		assert.NotNil(t, token)
// 	})
// }
