package logger

import (
	"context"
	"log"
	"retail_workflow/internal/shared/environment"
	"retail_workflow/internal/shared/errors"
	"retail_workflow/internal/shared/typings"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/redis/go-redis/v9/maintnotifications"
)

var client *redis.Client
var ctx context.Context

func connect() error {
	if client != nil && ctx != nil {
		return nil
	}

	ctx = context.TODO()
	client = redis.NewClient(&redis.Options{
		Addr:     environment.GetEnv(environment.LOGGING_ADDR),
		Username: environment.GetEnv(environment.LOGGING_USERNAME),
		Password: environment.GetEnv(environment.LOGGING_PASSWORD),
		DB:       0,
		MaintNotificationsConfig: &maintnotifications.Config{
			Mode: maintnotifications.ModeDisabled,
		},
	})
	cmd := client.Ping(ctx)

	if cmd.Val() != "PONG" {
		return errors.Error500.ErrRedisConnectionFailure
	}

	return nil
}

func Info(content string) {
	if err := connect(); err != nil {
		log.Println(err)
		return
	}

	key := time.Now().Format(time.RFC3339)
	cmd := client.Set(ctx, key, content, 0)

	if err := cmd.Err(); err != nil {
		log.Println(err)
		return
	}
}

func Error(errorResult typings.ErrorResult) {
	if err := connect(); err != nil {
		log.Println(err)
		return
	}

	key := time.Now().Format(time.RFC3339)
	cmd := client.JSONSet(ctx, key, "$", errorResult)

	if err := cmd.Err(); err != nil {
		log.Println(err)
		return
	}
}
