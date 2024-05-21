package model

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init CreateUser model", zap.String("journey", "createUser"))

	if err := ud.EncryptPassword(); err != nil {
		return rest_err.NewInternalServerError("Error trying to encrypt password")
	}

	return nil
}
