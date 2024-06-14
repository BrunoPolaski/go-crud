package service

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
	"go.uber.org/zap"
)

func (us *userDomainService) FindUserByIDService(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByID service", zap.String("journey", "findUserByID"))

	return us.repository.FindUserByIDRepository(id)
}

func (us *userDomainService) FindUserByEmailService(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail service", zap.String("journey", "findUserByEmail"))

	return us.repository.FindUserByEmailRepository(email)
}

func (us *userDomainService) FindUserByPasswordService(password string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByPassword service", zap.String("journey", "findUserByPassword"))

	return us.repository.FindUserByPasswordRepository(password)
}
