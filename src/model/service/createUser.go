package service

import (
	"fmt"

	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	"github.com/BrunoPolaski/go-crud/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init CreateUser model", zap.String("journey", "createUser"))

	if err := userDomain.EncryptPassword(); err != nil {
		return rest_err.NewInternalServerError("Error trying to encrypt password")
	}

	fmt.Println("UserDomain password: ", userDomain.GetPassword())

	return nil
}
