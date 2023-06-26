package repository

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/syamozipc/web_app/internal/database"
)

func GetTask(id string) (*sql.Row, error) {
	db, err, closer := database.SqlOpen()
	defer func() { _ = closer() }()
	if err != nil {
		return nil, err
	}

	row := db.QueryRow("SELECT * FROM tasks WHERE ID = $1", id)
	if err = row.Err(); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return row, nil
}

func ListTasks() (*sql.Rows, error) {
	db, err, closer := database.SqlOpen()
	defer func() { _ = closer() }()
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func CreateTask(title string) (string, error) {
	db, err, closer := database.SqlOpen()
	defer func() { _ = closer() }()
	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var id string
	row := db.QueryRow("INSERT INTO tasks (title) VALUES ($1) RETURNING id", title)
	if err = row.Err(); err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err = row.Scan(&id); err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return id, nil
}
