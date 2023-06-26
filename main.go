package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

func rootGet(c echo.Context) error {
	return c.String(http.StatusOK, "home")
}

type DomainTodo struct {
	ID        string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func SqlOpen() (*sql.DB, error) {
	return sql.Open("postgres", "postgres://root:root@localhost:54320/web_app?sslmode=disable")
}

func ListTodos(c echo.Context) error {
	db, err := SqlOpen()
	defer func() { _ = db.Close() }()
	if err != nil {
		return err
	}

	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		return err
	}

	var todos []DomainTodo
	for rows.Next() {
		var t DomainTodo
		err := rows.Scan(&t.ID, &t.Title, &t.CreatedAt, &t.UpdatedAt)
		if err != nil {
			return err
		}
		todos = append(todos, t)
	}

	var res = ResTodoList{
		List: ToTodoList(todos),
	}

	return c.JSON(http.StatusOK, res)
}

type ResTodoList struct {
	List []ResTodo `json:"todo_list"`
}

type ResTodo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

const JPTimeLayout = "2006/01/02 15:04:05"

func ToTodo(todo DomainTodo) ResTodo {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err.Error())
	}
	time.Local = loc

	res := ResTodo{
		ID:        todo.ID,
		Title:     todo.Title,
		CreatedAt: todo.CreatedAt.In(loc).Format(JPTimeLayout),
		UpdatedAt: todo.UpdatedAt.In(loc).Format(JPTimeLayout),
	}

	return res
}

func ToTodoList(todos []DomainTodo) []ResTodo {

	res := make([]ResTodo, 0, len(todos))
	for _, v := range todos {
		res = append(res, ToTodo(v))
	}

	return res
}

type ReqCreateTodo struct {
	Title string `json:"title" validate:"required"`
}

func CreateTodo(c echo.Context) error {
	var req ReqCreateTodo
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err = c.Validate(&req); err != nil {
		return err
	}

	db, err := SqlOpen()
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

	var res DomainTodo
	if err = row.Scan(&res.ID, &res.Title, &res.CreatedAt, &res.UpdatedAt); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, ToTodo(res))
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/", rootGet)
	e.GET("/todos", ListTodos)
	e.POST("/todos", CreateTodo)

	e.Logger.Fatal(e.Start(":8082"))
}
