package user

import "retail_workflow/internal/shared/valueobjects"

type UserEntity struct {
	Id         valueobjects.EntityId
	Password   valueobjects.Password
	Name       string
	EmployeeId string
}

func (u *UserEntity) New(dto UserCreationDTO) error {
	var id valueobjects.EntityId
	var password valueobjects.Password

	id.Generate()

	if err := password.Hash(dto.Password); err != nil {
		return err
	}

	u.Id = id
	u.Password = password
	u.Name = dto.Name
	u.EmployeeId = dto.EmployeeId

	return nil
}
