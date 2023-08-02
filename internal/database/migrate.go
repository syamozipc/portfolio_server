package database

import (
	"database/sql"
	"fmt"

	migrate "github.com/rubenv/sql-migrate"
	"github.com/syamozipc/web_app/internal/config"
)

func MigrateUp() error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	dsn := fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable", cfg.DB.Driver, cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		return err
	}
	fmt.Printf("Applied %d migrations!\n", n)

	return err
}
