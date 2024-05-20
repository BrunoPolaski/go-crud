package model

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (*UserDomain) CreateUser(UserDomain) *rest_err.RestErr {
	logger.Info("Init CreateUser model", zap.String("journey", "UserDomain"))

	return nil
}
