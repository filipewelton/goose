package repositories

import (
	"retail_workflow/internal/domain/user"
	"retail_workflow/internal/infrastructure/postgres"
	"retail_workflow/internal/persistence/models"
	"retail_workflow/internal/shared/errors"
	"retail_workflow/internal/shared/logger"
	"retail_workflow/internal/shared/typings"
)

type PostgresUserRepository struct{}

func (PostgresUserRepository) Insert(entity user.UserEntity) error {
	if err := postgres.Connect(); err != nil {
		return err
	}

	defer postgres.Disconnect()

	var model models.PostgresUserModel

	model.MapFromEntity(entity)

	tx := postgres.Client.Create(&model)

	if err := tx.Error; err != nil {
		return err
	}

	return nil
}

func (PostgresUserRepository) FindByEmployeeId(employeeId string) (
	user.UserEntity,
	error,
) {
	var entity user.UserEntity

	if err := postgres.Connect(); err != nil {
		return entity, err
	}

	defer postgres.Disconnect()

	var model models.PostgresUserModel
	var tx = postgres.Client.Find(&model, "employee_id = ?", employeeId)

	if err := tx.Error; err != nil {
		logger.Error(typings.ErrorResult{
			Code:    404,
			Context: "User not found by employee ID",
			Reason:  err.Error(),
		})

		return entity, errors.Err404.ErrUserNotFound
	}

	return model.MapToEntity(), nil
}

func (PostgresUserRepository) FindById(id string) (
	user.UserEntity,
	error,
) {
	var entity user.UserEntity

	if err := postgres.Connect(); err != nil {
		return entity, err
	}

	defer postgres.Disconnect()

	var model models.PostgresUserModel
	var tx = postgres.Client.Find(&model, "id = ?", id)

	if err := tx.Error; err != nil {
		logger.Error(typings.ErrorResult{
			Code:    404,
			Context: "User not found by ID",
			Reason:  err.Error(),
		})

		return entity, errors.Err404.ErrUserNotFound
	}

	return model.MapToEntity(), nil
}

func (PostgresUserRepository) DeleteByEmployeeId(employeeId string) error {
	if err := postgres.Connect(); err != nil {
		return err
	}

	defer postgres.Disconnect()

	var model models.PostgresUserModel
	var tx = postgres.Client.Delete(&model, "employee_id = ?", employeeId)

	if tx.Error != nil {
		return errors.Error500.ErrUserDeletionFailedInPostgreSQL
	}

	return nil
}
