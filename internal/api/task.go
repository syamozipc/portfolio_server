package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/syamozipc/web_app/internal/api/request"
	"github.com/syamozipc/web_app/internal/api/response"
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

	task, err := repository.GetTask(req.ID)
	if err != nil {
		return err
	}

	res := response.ToTask(*task)

	return c.JSON(http.StatusOK, res)
}

func ListTasks(c echo.Context) error {
	tasks, err := repository.ListTasks()
	if err != nil {
		return err
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

	task, err := repository.GetTask(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, response.ToTask(*task))
}

func UpdateTask(c echo.Context) error {
	var req request.UpdateTask
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	if err := repository.UpdateTask(req.ID, req.Title); err != nil {
		return err
	}

	task, err := repository.GetTask(req.ID)
	if err != nil {
		return err
	}

	res := response.ToTask(*task)

	return c.JSON(http.StatusOK, res)
}

func DeteleTask(c echo.Context) error {
	var req request.DeteleTask
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	if err := repository.DeleteTask(req.ID); err != nil {
		return err
	}

	return c.JSON(http.StatusNoContent, nil)
}
