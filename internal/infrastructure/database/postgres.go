package database

import (
	"os"
	"retail_flow/internal/shared/exceptions"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Postgres struct {
	Client *gorm.DB
}

func (p *Postgres) Connect() error {
	dsn := os.Getenv("DATABASE_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(nil, logger.Config{}),
	})

	if err != nil {
		return exceptions.InternalExceptions.ErrPostgresConnectionFailure
	}

	sql, err := db.DB()

	if err != nil {
		return exceptions.InternalExceptions.ErrConnectionToPostgresNotEstablished
	} else if err = sql.Ping(); err != nil {
		return exceptions.InternalExceptions.ErrConnectionToPostgresNotEstablished
	}

	p.Client = db
	return nil
}

func (p *Postgres) Disconnect() error {
	if p.Client == nil {
		return exceptions.InternalExceptions.ErrConnectionToPostgresNotEstablished
	}

	sql, err := p.Client.DB()

	if err != nil {
		return exceptions.InternalExceptions.ErrConnectionToPostgresNotEstablished
	} else if err = sql.Close(); err != nil {
		return exceptions.InternalExceptions.ErrConnectionToPostgresNotEstablished
	}

	p.Client = nil
	return nil
}
