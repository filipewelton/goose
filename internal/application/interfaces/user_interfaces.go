package interfaces

import "retail_flow/internal/domain/entities"

type UserRepository interface {
	Insert(entities.UserEntity) error
	FindByCardNumber(string) (entities.UserEntity, error)
	DeleteByID(string) error
}
