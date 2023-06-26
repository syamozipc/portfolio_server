package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

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

func ListTodos(c echo.Context) error {
	db, err := sql.Open("postgres", "postgres://root:root@localhost:54320/web_app?sslmode=disable")
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

func main() {
	e := echo.New()
	e.GET("/", rootGet)
	e.GET("/todos", ListTodos)

	e.Logger.Fatal(e.Start(":8082"))
}
