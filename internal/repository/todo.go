package repository

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/syamozipc/web_app/internal/database"
)

func GetTask(id string) (*sql.Row, error) {
	db, err := database.Open()
	defer func() { _ = db.Close() }()
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
	db, err := database.Open()
	defer func() { _ = db.Close() }()
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
	db, err := database.Open()
	defer func() { _ = db.Close() }()
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

func UpdateTask(id, title string) error {
	db, err := database.Open()
	defer func() { _ = db.Close() }()
	if err != nil {
		return err
	}

	row := db.QueryRow("UPDATE tasks SET title = $1, updated_at = $2 WHERE id = $3", title, time.Now(), id)
	if row.Err() != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func DeleteTask(id string) error {
	db, err := database.Open()
	defer func() { _ = db.Close() }()
	if err != nil {
		return err
	}

	row := db.QueryRow("DELETE FROM tasks WHERE id = $1", id)
	if row.Err() != nil {
		return row.Err()
	}

	return nil
}
