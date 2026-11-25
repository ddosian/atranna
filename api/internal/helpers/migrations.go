package helpers

import (
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"

	"github.com/ddosian/atranna/api/migrations"
)

func RunMigrations(conn *sql.DB) error {
	goose.SetBaseFS(migrations.FS)
	defer goose.SetBaseFS(nil)

	if err := goose.SetDialect("sqlite3"); err != nil {
		return fmt.Errorf("set goose dialect: %w", err)
	}
	if err := goose.Up(conn, "."); err != nil {
		return fmt.Errorf("run migrations: %w", err)
	}
	return nil
}
