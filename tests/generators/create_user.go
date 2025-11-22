package generators

import (
	"retail_workflow/internal/domain/user"
	"retail_workflow/tests/utils"

	"syreclabs.com/go/faker"
)

func CreateUser(repository user.UserRepository) user.UserEntity {
	var entity user.UserEntity

	entity.New(user.UserCreationDTO{
		Name:       faker.Name().Name(),
		EmployeeId: faker.Number().Number(7),
		Password:   utils.GeneratePassword(),
	})

	if err := repository.Insert(entity); err != nil {
		panic(err)
	}

	return entity
}
