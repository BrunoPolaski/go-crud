package service

import (
	"github.com/BrunoPolaski/go-crud/src/configuration/rest_err"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
)

func (*userDomainService) UpdateUser(userDomain model.UserDomainInterface, userId string) *rest_err.RestErr {
	return nil
}
