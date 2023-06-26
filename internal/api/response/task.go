package response

import (
	"log"
	"time"

	"github.com/syamozipc/web_app/internal/constant"
	"github.com/syamozipc/web_app/internal/model"
)

type TaskList struct {
	List []Task `json:"task_list"`
}

type Task struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ToTask(task model.Task) Task {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err.Error())
	}
	time.Local = loc

	res := Task{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt.In(loc).Format(constant.JPTimeLayout),
		UpdatedAt: task.UpdatedAt.In(loc).Format(constant.JPTimeLayout),
	}

	return res
}

func ToTaskList(tasks []model.Task) []Task {

	res := make([]Task, 0, len(tasks))
	for _, v := range tasks {
		res = append(res, ToTask(v))
	}

	return res
}
