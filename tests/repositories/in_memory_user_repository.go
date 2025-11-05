package repositories

import (
	"retail_flow/internal/domain/entities"
	"retail_flow/internal/shared/exceptions"
)

type InMemoryUserRepository struct{}

var userDB = map[string]entities.UserEntity{}

func (InMemoryUserRepository) Insert(user entities.UserEntity) error {
	id := user.ID.GetValue()
	userDB[id] = user

	return nil
}

func (InMemoryUserRepository) FindByCardNumber(
	cardNumber string,
) (entities.UserEntity, error) {
	for _, v := range userDB {
		if v.CardNumber == cardNumber {
			return v, nil
		}
	}

	return entities.UserEntity{}, exceptions.NotFoundExceptions.ErrUserNotFound
}

func (InMemoryUserRepository) DeleteByID(id string) error {
	searchResult := userDB[id]

	if searchResult == (entities.UserEntity{}) {
		return exceptions.NotFoundExceptions.ErrUserNotFound
	}

	delete(userDB, id)
	return nil
}
