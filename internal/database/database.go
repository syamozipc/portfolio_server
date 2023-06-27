package database

import "database/sql"

func Open() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://root:root@localhost:54320/web_app?sslmode=disable")
	if err != nil {
		return nil, err
	}

	return db, nil
}
