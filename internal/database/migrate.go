package database

import (
	"database/sql"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
)

func MigrateUp() error {
	migrations := &migrate.FileMigrationSource{
		Dir: "internal/migrations",
	}

	db, err := sql.Open("postgres", "postgres://root:root@localhost:54320/web_app?sslmode=disable")
	if err != nil {
		return err
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}
	fmt.Printf("Applied %d migrations!\n", n)

	return err
}
