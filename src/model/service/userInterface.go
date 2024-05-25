package service

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/BrunoPolaski/go-crud/src/model"
)

type userDomainService struct{}

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type UserDomainService interface {
	CreateUserService(model.UserDomainInterface) *rest_err.RestErr
	UpdateUserService(model.UserDomainInterface, string) *rest_err.RestErr
	FindUserService(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUserService(string) *rest_err.RestErr
}
