package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://root:root@localhost:54320/web_app?sslmode=disable")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
