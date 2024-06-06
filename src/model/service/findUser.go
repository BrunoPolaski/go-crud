package service

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
	"go.uber.org/zap"
)

func (us *userDomainService) FindUserByID(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByID service", zap.String("journey", "findUserByID"))

	return us.repository.FindUserByID(id)
}

func (us *userDomainService) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail service", zap.String("journey", "findUserByEmail"))

	return us.repository.FindUserByEmail(email)
}
