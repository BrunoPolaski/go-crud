package service

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/BrunoPolaski/go-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateUser service", zap.String("journey", "createUser"))

	if err := userDomain.EncryptPassword(); err != nil {
		return nil, rest_err.NewInternalServerError("Error trying to encrypt password")
	}

	if userDomainRepository, err := ud.repository.CreateUser(userDomain); err != nil {
		return nil, err
	} else {
		return userDomainRepository, nil
	}
}
