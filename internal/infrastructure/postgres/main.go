package postgres

import (
	"retail_workflow/internal/shared/environment"
	"retail_workflow/internal/shared/errors"
	"retail_workflow/internal/shared/logger"
	"retail_workflow/internal/shared/typings"

	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Client *gorm.DB

func Connect() error {
	dsn := environment.GetEnv(environment.POSTGRES_DSN)
	db, err := gorm.Open(pg.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Error(typings.ErrorResult{
			Code:    500,
			Context: "PostreSQL connection",
			Reason:  err.Error(),
		})

		return errors.Error500.ErrPostgresConnectionFailure
	}

	sql, err := db.DB()

	if err != nil {
		logger.Error(typings.ErrorResult{
			Code:    500,
			Context: "PostreSQL connection",
			Reason:  err.Error(),
		})

		return errors.Error500.ErrPostgresConnectionFailure
	} else if err = sql.Ping(); err != nil {
		logger.Error(typings.ErrorResult{
			Code:    500,
			Context: "PostreSQL connection",
			Reason:  err.Error(),
		})

		return errors.Error500.ErrPostgresConnectionFailure
	}

	Client = db
	return nil
}

func Disconnect() {
	if Client == nil {
		return
	}

	sql, err := Client.DB()

	if err != nil {
		logger.Error(typings.ErrorResult{
			Code:    500,
			Context: "PostreSQL disconnection",
			Reason:  err.Error(),
		})

		return
	}

	sql.Close()
}
