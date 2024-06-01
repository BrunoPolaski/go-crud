package converter

import (
	"github.com/BrunoPolaski/go-crud/src/model"
	"github.com/BrunoPolaski/go-crud/src/model/repository/entity"
)

func ConvertEntityToDomain(entity *entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)

	domain.SetId(entity.Id)

	return domain
}
