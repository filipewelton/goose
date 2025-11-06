package repositories

import (
	"retail_flow/internal/infrastructure/cache"
	"retail_flow/internal/shared/exceptions"
	"time"
)

type RedisWhitelistRepository struct{}

func (RedisWhitelistRepository) Insert(cardNumber string) error {
	var redis cache.Redis

	if err := redis.Connect(); err != nil {
		return err
	}

	cmd := redis.Client.Set(*redis.Ctx, cardNumber, "OK", time.Hour*24*7)

	if err := cmd.Err(); err != nil {
		return exceptions.InternalExceptions.ErrRedisInsertionFailed
	}

	return redis.Disconnect()
}

func (RedisWhitelistRepository) Has(cardNumber string) (bool, error) {
	var redis cache.Redis

	if err := redis.Connect(); err != nil {
		return false, err
	}

	cmd := redis.Client.Get(*redis.Ctx, cardNumber)

	if err := cmd.Err(); err != nil {
		return false, exceptions.NotFoundExceptions.ErrPreRegisterNotFound
	} else if err = redis.Disconnect(); err != nil {
		return false, err
	}

	return true, nil
}

func (RedisWhitelistRepository) Delete(cardNumber string) error {
	var redis cache.Redis

	if err := redis.Connect(); err != nil {
		return err
	}

	cmd := redis.Client.Del(*redis.Ctx, cardNumber)

	if err := cmd.Err(); err != nil {
		return exceptions.InternalExceptions.ErrRedisDeletionFailed
	} else if err = redis.Disconnect(); err != nil {
		return err
	}

	return nil
}
