package mocks

import (
	"retail_workflow/internal/domain/user"
	"retail_workflow/internal/shared/errors"
)

type InMemoryUserRepository struct{}

var inMemoryUser = map[string]user.UserEntity{}

func (InMemoryUserRepository) Insert(user user.UserEntity) error {
	inMemoryUser[user.Id.Get()] = user
	return nil
}

func (InMemoryUserRepository) FindByEmployeeId(employeeId string) (
	user.UserEntity,
	error,
) {
	for _, v := range inMemoryUser {
		if v.EmployeeId == employeeId {
			return v, nil
		}
	}

	return user.UserEntity{}, errors.Err404.ErrUserNotFound
}

func (InMemoryUserRepository) FindById(id string) (user.UserEntity, error) {
	for k, v := range inMemoryUser {
		if k == id {
			return v, nil
		}
	}

	return user.UserEntity{}, errors.Err404.ErrUserNotFound
}

func (InMemoryUserRepository) DeleteByEmployeeId(id string) error {
	delete(inMemoryUser, id)
	return nil
}
