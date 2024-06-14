package service

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
	"go.uber.org/zap"
)

func (us *userDomainService) LoginUserService(userDomain userDomainService) (model.UserDomainInterface, string, *rest_err.RestErr) {
	logger.Info("Init LoginUser service", zap.String("journey", "loginUser"))

	_, err := us.repository.FindUserByEmailRepository(userRequest.Email)
	if err != nil {
		return err
	}

	_, err = us.repository.FindUserByPasswordRepository(userRequest.Password)
	if err != nil {
		return err
	}

	return nil
}
