package service

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
	"go.uber.org/zap"
)

func (us *userDomainService) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUser service", zap.String("journey", "findUser"))

	if user, err := us.repository.FindUserByEmail(email); err != nil {
		return nil, err
	} else {
		return user, nil
	}
}
