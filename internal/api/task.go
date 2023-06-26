package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/syamozipc/web_app/internal/api/request"
	"github.com/syamozipc/web_app/internal/api/response"
	"github.com/syamozipc/web_app/internal/database"
	"github.com/syamozipc/web_app/internal/model"
)

func ListTasks(c echo.Context) error {
	db, err := database.SqlOpen()
	defer func() { _ = db.Close() }()
	if err != nil {
		return err
	}

	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		return err
	}

	var tasks []model.Task
	for rows.Next() {
		var t model.Task
		err := rows.Scan(&t.ID, &t.Title, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return err
		}
		tasks = append(tasks, t)
	}

	var res = response.TaskList{
		List: response.ToTaskList(tasks),
	}

	return c.JSON(http.StatusOK, res)
}

func CreateTask(c echo.Context) error {
	var req request.CreateTask
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(&req); err != nil {
		return err
	}

	db, err := database.SqlOpen()
	defer func() { _ = db.Close() }()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var id string
	row := db.QueryRow("INSERT INTO tasks (title) VALUES ($1) RETURNING id", req.Title)
	if err = row.Err(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err = row.Scan(&id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	row = db.QueryRow("SELECT * FROM tasks WHERE ID = $1", id)
	if err = row.Err(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var res model.Task
	if err = row.Scan(&res.ID, &res.Title, &res.CreatedAt, &res.UpdatedAt); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, response.ToTask(res))
}
