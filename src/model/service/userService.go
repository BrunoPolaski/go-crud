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
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(model.UserDomainInterface, string) *rest_err.RestErr
	FindUserByEmail(string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByID(string) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
