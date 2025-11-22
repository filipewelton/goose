package redis

import (
	"context"
	"retail_workflow/internal/shared/environment"
	"retail_workflow/internal/shared/errors"

	r "github.com/redis/go-redis/v9"
	"github.com/redis/go-redis/v9/maintnotifications"
)

var Client *r.Client
var Ctx context.Context

func Connect() error {
	ctx := context.Background()
	rdb := r.NewClient(&r.Options{
		Addr:     environment.GetEnv(environment.CACHE_ADDR),
		Username: environment.GetEnv(environment.CACHE_USERNAME),
		Password: environment.GetEnv(environment.CACHE_PASSWORD),
		DB:       0,
		MaintNotificationsConfig: &maintnotifications.Config{
			Mode: maintnotifications.ModeDisabled,
		},
	})
	cmd := rdb.Ping(ctx)

	if cmd.Val() != "PONG" {
		return errors.Error500.ErrRedisConnectionFailure
	}

	Client = rdb
	Ctx = ctx
	return nil
}

func Disconnect() {
	if Client == nil || Ctx == nil {
		return
	}

	Client.Close()
	Ctx.Done()

	Client = nil
	Ctx = nil
}
