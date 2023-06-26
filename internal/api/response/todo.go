package response

import (
	"log"
	"time"

	"github.com/syamozipc/web_app/internal/constant"
	"github.com/syamozipc/web_app/internal/model"
)

type TodoList struct {
	List []Todo `json:"todo_list"`
}

type Todo struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ToTodo(todo model.Todo) Todo {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err.Error())
	}
	time.Local = loc

	res := Todo{
		ID:        todo.ID,
		Title:     todo.Title,
		CreatedAt: todo.CreatedAt.In(loc).Format(constant.JPTimeLayout),
		UpdatedAt: todo.UpdatedAt.In(loc).Format(constant.JPTimeLayout),
	}

	return res
}

func ToTodoList(todos []model.Todo) []Todo {

	res := make([]Todo, 0, len(todos))
	for _, v := range todos {
		res = append(res, ToTodo(v))
	}

	return res
}
