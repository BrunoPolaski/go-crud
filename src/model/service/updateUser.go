package service

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUserService(userDomain model.UserDomainInterface, id string) *rest_err.RestErr {
	logger.Info("Init UpdateUser service", zap.String("journey", "UpdateUser"))

	if err := ud.repository.UpdateUserRepository(userDomain, id); err != nil {
		return err
	} else {
		return nil
	}
}
