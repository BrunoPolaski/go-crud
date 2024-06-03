package converter

import (
	"github.com/BrunoPolaski/go-crud/src/model/repository/entity"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
)

func ConvertDomainToEntity(domain model.UserDomainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}
