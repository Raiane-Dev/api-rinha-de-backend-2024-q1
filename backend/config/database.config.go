package config

import (
	"fmt"
	"rinha_api/backend/util/logger"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	migrate "github.com/rubenv/sql-migrate"
)

var (
	DatabaseInstance *sqlx.DB
)

func ConnectInstance() (err error) {
	DatabaseInstance, err = sqlx.Open("sqlite3", "/data/rinha_api.sqlite?_journal=WAL&_timeout=5000&_fk=true")
	if err != nil {
		return
	}
	if err = DatabaseInstance.Ping(); err != nil {
		return
	}

	return
}

func ExecMigration() {

	migrations := &migrate.FileMigrationSource{
		Dir: "/data/schemas/",
	}

	n, err := migrate.Exec(DatabaseInstance.DB, "sqlite3", migrations, migrate.Up)
	if err != nil {
		logger.Error("Error applying migrations: %v", err)
	}

	if n != 0 {
		logger.Info(fmt.Sprintf("Applied %d migration", n))
	}

}
