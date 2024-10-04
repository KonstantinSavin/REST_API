package sqldb

import (
	"database/sql"

	"github.com/pressly/goose/v3"
)

var migrationPath = "/migrations"

// MigrationsUp applies all migrations.
func MigrationsUp(db *sql.DB) error {
	if err := goose.SetDialect("pgx"); err != nil {
		return err
	}

	return goose.Up(db, migrationPath)
}

// MigrationsDown rolls back migrations.
func MigrationsDown(db *sql.DB) error {
	if err := goose.SetDialect("pgx"); err != nil {
		return err
	}

	return goose.Down(db, migrationPath)
}
