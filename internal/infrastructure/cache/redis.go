package cache

import (
	"context"
	"os"
	"retail_flow/internal/shared/exceptions"

	"github.com/redis/go-redis/v9"
	"github.com/redis/go-redis/v9/maintnotifications"
)

type Redis struct {
	Ctx    *context.Context
	Client *redis.Client
}

func (r *Redis) Connect() error {
	ctx := context.Background()
	user := os.Getenv("REDIS_USER")
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_DSN"),
		Username: user,
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
		MaintNotificationsConfig: &maintnotifications.Config{
			Mode: maintnotifications.ModeDisabled,
		},
	})
	err := client.Ping(ctx).Err()

	if err != nil {
		return exceptions.InternalExceptions.ErrRedisConnectionFailure
	}

	r.Ctx = &ctx
	r.Client = client

	return nil
}

func (r *Redis) Disconnect() error {
	if err := r.Client.Close(); err != nil {
		return exceptions.InternalExceptions.ErrRedisDisconnectionFailure
	}

	r.Client = nil
	r.Ctx = nil
	return nil
}
