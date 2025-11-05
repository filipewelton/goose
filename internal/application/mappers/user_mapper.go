package mappers

import (
	"retail_flow/internal/domain/entities"
	"retail_flow/internal/shared/typings"
)

func MapUserEntity(user entities.UserEntity) typings.UserEntityMapped {
	return typings.UserEntityMapped{
		ID:         user.ID.GetValue(),
		CardNumber: user.CardNumber,
		Name:       user.Name,
		Supervisor: user.Supervisor,
	}
}
