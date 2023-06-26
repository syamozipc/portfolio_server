package database

import "database/sql"

func SqlOpen() (*sql.DB, error) {
	return sql.Open("postgres", "postgres://root:root@localhost:54320/web_app?sslmode=disable")
}
