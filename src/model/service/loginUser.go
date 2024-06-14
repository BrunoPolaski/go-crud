package service

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
	"go.uber.org/zap"
)

func (us *userDomainService) LoginUserService(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init LoginUser service", zap.String("journey", "loginUser"))

	user, err := us.FindUserByEmailService(userDomain.GetEmail())
	if err != nil {
		return nil, err
	}

	if err := user.ComparePassword(userDomain.GetPassword()); err != nil {
		return nil, rest_err.NewUnauthorizedError("Invalid credentials")
	}

	return user, nil
}
