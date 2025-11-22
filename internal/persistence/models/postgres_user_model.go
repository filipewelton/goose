package models

import (
	"retail_workflow/internal/domain/user"
	"retail_workflow/internal/shared/valueobjects"
)

type PostgresUserModel struct {
	Id         string `gorm:"column:id;primaryKey;not null"`
	EmployeeId string `gorm:"column:employee_id"`
	Name       string
	Password   string
}

func (PostgresUserModel) TableName() string {
	return "users"
}

func (p *PostgresUserModel) MapFromEntity(entity user.UserEntity) {
	p.Id = entity.Id.Get()
	p.Password = entity.Password.Get()
	p.EmployeeId = entity.EmployeeId
	p.Name = entity.Name
}

func (p *PostgresUserModel) MapToEntity() user.UserEntity {
	var id valueobjects.EntityId
	var password valueobjects.Password

	id.Set(p.Id)
	password.Set(p.Password)

	return user.UserEntity{
		Id:         id,
		Password:   password,
		Name:       p.Name,
		EmployeeId: p.EmployeeId,
	}
}
