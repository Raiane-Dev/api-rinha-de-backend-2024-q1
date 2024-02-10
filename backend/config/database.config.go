package config

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	migrate "github.com/rubenv/sql-migrate"
)

var (
	DatabaseInstance *sqlx.DB
)

func ConnectInstance() (err error) {
	DatabaseInstance, err = sqlx.Open("sqlite3", "/data/rinha_api.sqlite")
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
		log.Fatalf("Error applying migrations: %v", err)
	}

	if n != 0 {
		log.Fatal(fmt.Sprintf("Applied %d migration", n))
	}

}
