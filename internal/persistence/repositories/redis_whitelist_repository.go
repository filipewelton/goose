package repositories

import (
	"retail_workflow/internal/infrastructure/redis"
	"retail_workflow/internal/shared/errors"
	"retail_workflow/internal/shared/logger"
	"retail_workflow/internal/shared/typings"
)

type RedisWhitelistRepository struct{}

const ok = "OK"

func (RedisWhitelistRepository) Insert(employeeId string) error {
	if err := redis.Connect(); err != nil {
		return err
	}

	defer redis.Disconnect()

	cmd := redis.Client.Set(redis.Ctx, employeeId, ok, 0)

	if err := cmd.Err(); err != nil {
		logger.Error(typings.ErrorResult{
			Code:    500,
			Context: "Employee ID insertion in whitelist",
			Reason:  err.Error(),
		})

		return errors.Error500.ErrKeyInsertionFailedInRedis
	}

	return nil
}

func (RedisWhitelistRepository) Has(employeeId string) (bool, error) {
	if err := redis.Connect(); err != nil {
		return false, err
	}

	defer redis.Disconnect()

	cmd := redis.Client.Get(redis.Ctx, employeeId)

	if err := cmd.Err(); err != nil {
		logger.Error(typings.ErrorResult{
			Code:    404,
			Context: "Employee ID registration verification",
			Reason:  err.Error(),
		})

		return false, errors.Err404.ErrEmployeeIdNotFound
	} else if cmd.Val() != ok {
		return false, errors.Err404.ErrEmployeeIdNotFound
	}

	return true, nil
}

func (RedisWhitelistRepository) Delete(employeeId string) error {
	if err := redis.Connect(); err != nil {
		return err
	}

	defer redis.Disconnect()

	cmd := redis.Client.Del(redis.Ctx, employeeId)

	if err := cmd.Err(); err != nil {
		logger.Error(typings.ErrorResult{
			Code:    404,
			Context: "Employee ID deletion",
			Reason:  err.Error(),
		})

		return errors.Err404.ErrEmployeeIdNotFound
	}

	return nil
}
