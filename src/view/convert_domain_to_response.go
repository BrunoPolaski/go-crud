package view

import (
	"github.com/BrunoPolaski/go-crud/src/controller/model/response"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetID(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
