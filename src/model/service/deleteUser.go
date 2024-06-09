package service

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (us *userDomainService) DeleteUser(id string) *rest_err.RestErr {
	logger.Info("Init DeleteUser service", zap.String("journey", "DeleteUser"))

	return us.repository.DeleteUserRepository(id)
}
