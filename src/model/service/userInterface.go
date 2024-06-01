package service

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/restErr"
	"github.com/BrunoPolaski/go-crud/src/model"
	"github.com/BrunoPolaski/go-crud/src/model/repository"
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
	CreateUser(model.UserDomainInterface) *restErr.RestErr
	UpdateUser(model.UserDomainInterface, string) *restErr.RestErr
	FindUser(string) (*model.UserDomainInterface, *restErr.RestErr)
	DeleteUser(string) *restErr.RestErr
}
