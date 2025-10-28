package lib

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func ConnectDatabase(driver, dsn string) {
	db, err := sql.Open(driver, dsn)

	if err != nil {
		ThrowError("Failed to connect to database", err)
	}

	DB = db
}

func DisconnectDatabase() {
	if err := DB.Close(); err != nil {
		ThrowError("Failed to disconnect from database", err)
	}
}
