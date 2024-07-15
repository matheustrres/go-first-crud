package converter

import (
	"github.com/matheustrres/go-first-crud/src/model"
	"github.com/matheustrres/go-first-crud/src/model/repository/entity"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) model.UserDomainInterface {
	userDomain := model.NewUserDomain(
		entity.Name,
		entity.Email,
		entity.Password,
		entity.Age,
	)

	userDomain.SetID(entity.ID.String())

	return userDomain
}
