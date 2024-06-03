package converter

import (
	"github.com/BrunoPolaski/go-crud/src/model/repository/entity"
	model "github.com/BrunoPolaski/go-crud/src/model/user"
)

func ConvertEntityToDomain(entity *entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)

	domain.SetID(entity.ID.Hex())

	return domain
}
