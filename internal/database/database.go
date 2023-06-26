package database

import "database/sql"

func SqlOpen() (*sql.DB, error, func() error) {
	db, err := sql.Open("postgres", "postgres://root:root@localhost:54320/web_app?sslmode=disable")
	if err != nil {
		return nil, err, db.Close
	}

	return db, nil, db.Close
}
