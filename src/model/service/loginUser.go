package service

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/logger"
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
	"go.uber.org/zap"
)

func (us *userDomainService) LoginUserService(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr) {
	logger.Info("Init LoginUser service", zap.String("journey", "loginUser"))

	users, err := us.FindAllUsersService(userDomain.GetEmail())
	if err != nil {
		return nil, "", err
	}

	if err := users[0].ComparePassword(userDomain.GetPassword()); err != nil {
		return nil, "", rest_err.NewUnauthorizedError("Invalid credentials")
	}

	token, err := users[0].GenerateToken()
	if err != nil {
		return nil, "", err
	}

	return users[0], token, nil
}
