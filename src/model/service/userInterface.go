package service

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/BrunoPolaski/go-crud/src/model"
	"github.com/BrunoPolaski/go-crud/src/model/repository"
)

type userDomainService struct{}

func NewUserDomainService(
	repository repository.UserRepository,
) UserDomainService {
	return &userDomainService{}
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_err.RestErr
	UpdateUser(model.UserDomainInterface, string) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
