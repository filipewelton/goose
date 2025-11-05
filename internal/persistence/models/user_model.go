package models

import (
	"retail_flow/internal/domain/entities"
	"retail_flow/internal/domain/valueobjects"
)

type UserModel struct {
	ID         string `gorm:"type:varchar(255);primaryKey;not null"`
	CardNumber string `gorm:"column:card_number;type:varchar(7);unique;not null"`
	Name       string `gorm:"type:varchar(50);not null"`
	Password   string `gorm:"type:varchar(255);not null"`
	Supervisor bool   `gorm:"type:bool;not null;default:false"`
}

func (u *UserModel) MapToModel(user entities.UserEntity) {
	u.ID = user.ID.GetValue()
	u.CardNumber = user.CardNumber
	u.Name = user.Name
	u.Password = user.Password.GetValue()
	u.Supervisor = user.Supervisor
}

func (u *UserModel) MapToEntity() entities.UserEntity {
	var id valueobjects.EntityID
	var password valueobjects.Password

	id.SetValue(u.ID)
	password.SetValue(u.Password)

	return entities.UserEntity{
		ID:         id,
		Password:   password,
		CardNumber: u.CardNumber,
		Name:       u.Name,
		Supervisor: u.Supervisor,
	}
}

func (u *UserModel) TableName() string {
	return "users"
}
