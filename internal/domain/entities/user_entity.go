package entities

import (
	"retail_flow/internal/application/dto"
	"retail_flow/internal/domain/valueobjects"
)

type UserEntity struct {
	ID         valueobjects.EntityID
	Password   valueobjects.Password
	CardNumber string
	Name       string
	Supervisor bool
}

func (u *UserEntity) New(dto dto.UserCreationDTO) error {
	var (
		id       valueobjects.EntityID
		password valueobjects.Password
	)

	if err := id.New(); err != nil {
		return err
	} else if err = password.Hash(dto.Password); err != nil {
		return err
	}

	u.ID = id
	u.Password = password
	u.CardNumber = dto.CardNumber
	u.Name = dto.Name
	u.Supervisor = false

	return nil
}
