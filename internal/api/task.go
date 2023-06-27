package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/syamozipc/web_app/internal/api/request"
	"github.com/syamozipc/web_app/internal/api/response"
	"github.com/syamozipc/web_app/internal/model"
	"github.com/syamozipc/web_app/internal/repository"
)

func GetTask(c echo.Context) error {
	var req request.GetTask

	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	row, err := repository.GetTask(req.ID)

	if err != nil {
		return err
	}

	var task model.Task
	if err := row.Scan(&task.ID, &task.Title, &task.CreatedAt, &task.UpdatedAt); err != nil {
		return err
	}

	res := response.ToTask(task)

	return c.JSON(http.StatusOK, res)
}

func ListTasks(c echo.Context) error {
	rows, err := repository.ListTasks()
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

	id, err := repository.CreateTask(req.Title)
	if err != nil {
		return err
	}

	row, err := repository.GetTask(id)
	if err != nil {
		return err
	}

	var res model.Task
	if err = row.Scan(&res.ID, &res.Title, &res.CreatedAt, &res.UpdatedAt); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, response.ToTask(res))
}
