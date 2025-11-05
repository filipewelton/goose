package repositories

import (
	"retail_flow/internal/domain/entities"
	"retail_flow/internal/infrastructure/database"
	"retail_flow/internal/persistence/models"
	"retail_flow/internal/shared/exceptions"
)

type PostgresUserRepository struct{}

var postgres database.Postgres

func (PostgresUserRepository) Insert(user entities.UserEntity) error {

	if err := postgres.Connect(); err != nil {
		return err
	}

	var model models.UserModel

	model.MapToModel(user)

	result := postgres.Client.Create(&model)

	if result.Error != nil {
		return exceptions.InternalExceptions.ErrFailedToInsertUserIntoDatabase
	} else if err := postgres.Disconnect(); err != nil {
		return err
	}

	return nil
}

func (PostgresUserRepository) FindByCardNumber(
	cardNumber string,
) (entities.UserEntity, error) {
	var entity entities.UserEntity
	var model models.UserModel

	if err := postgres.Connect(); err != nil {
		return entity, err
	}

	result := postgres.Client.Find(&model, "card_number = ?", cardNumber)

	if result.Error != nil {
		return entity, exceptions.NotFoundExceptions.ErrUserNotFound
	} else if err := postgres.Disconnect(); err != nil {
		return entity, err
	}

	entity = model.MapToEntity()

	return entity, nil
}

func (PostgresUserRepository) DeleteByID(id string) error {
	if err := postgres.Connect(); err != nil {
		return err
	}

	var model models.UserModel
	var result = postgres.Client.Delete(&model, "id = ?", id)

	if result.Error != nil {
		return exceptions.NotFoundExceptions.ErrUserNotFound
	} else if err := postgres.Disconnect(); err != nil {
		return err
	}

	return nil
}
