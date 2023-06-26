package api

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/syamozipc/web_app/internal/api/request"
	"github.com/syamozipc/web_app/internal/api/response"
	"github.com/syamozipc/web_app/internal/database"
	"github.com/syamozipc/web_app/internal/model"
)

func ListTodos(c echo.Context) error {
	db, err := database.SqlOpen()
	defer func() { _ = db.Close() }()
	if err != nil {
		return err
	}

	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		return err
	}

	var todos []model.Todo
	for rows.Next() {
		var t model.Todo
		err := rows.Scan(&t.ID, &t.Title, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return err
		}
		todos = append(todos, t)
	}

	var res = response.TodoList{
		List: response.ToTodoList(todos),
	}

	return c.JSON(http.StatusOK, res)
}

func CreateTodo(c echo.Context) error {
	var req request.CreateTodo
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
	row := db.QueryRow("INSERT INTO todos (title) VALUES ($1) RETURNING id", req.Title)
	if err = row.Err(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err = row.Scan(&id); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	row = db.QueryRow("SELECT * FROM todos WHERE ID = $1", id)
	if err = row.Err(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var res model.Todo
	if err = row.Scan(&res.ID, &res.Title, &res.CreatedAt, &res.UpdatedAt); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, response.ToTodo(res))
}
