package service

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (us *userDomainService) DeleteUser(id string) *rest_err.RestErr {
	logger.Info("Init DeleteUser service", zap.String("journey", "DeleteUser"))

	err := us.repository.DeleteUserRepository(id)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "DeleteUser"))
		return err
	}

	logger.Info("Successfully executed DeleteUser service", zap.String("journey", "DeleteUser"))
	return nil
}
