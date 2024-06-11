package service

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserService(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init CreateUser service", zap.String("journey", "createUser"))

	if userDomain, _ := ud.repository.FindUserByEmailRepository(userDomain.GetEmail()); userDomain != nil {
		logger.Info("User already exists", zap.String("journey", "createUser"))
		return nil, rest_err.NewBadRequestError("User already exists")
	}

	if err := userDomain.EncryptPassword(); err != nil {
		return nil, rest_err.NewInternalServerError("Error trying to encrypt password")
	}

	return ud.repository.CreateUserRepository(userDomain)
}
