package service

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	repository "github.com/BrunoPolaski/go-crud/src/model/repository/user"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
)

type userDomainService struct {
	repository repository.UserRepository
}

func NewUserDomainService(
	repository repository.UserRepository,
) UserDomainService {
	return &userDomainService{
		repository: repository,
	}
}

type UserDomainService interface {
	CreateUserService(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUserService(model.UserDomainInterface, string) *rest_err.RestErr
	FindUserByEmailService(string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByIDService(string) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUserService(string) *rest_err.RestErr
	LoginUserService(model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr)
}
