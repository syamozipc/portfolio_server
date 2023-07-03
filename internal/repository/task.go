package repository

import (
	"time"

	"github.com/google/uuid"

	"github.com/syamozipc/web_app/internal/database"
	"github.com/syamozipc/web_app/internal/model"
)

func GetTask(id string) (*model.Task, error) {
	db := database.Pool()

	var task model.Task
	if err := db.Debug().Where("id = ?", id).First(&task).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func ListTasks() ([]model.Task, error) {
	db := database.Pool()

	var tasks []model.Task
	if err := db.Debug().Find(&tasks).Error; err != nil {
		return nil, err
	}

	return tasks, nil
}

func CreateTask(title string) (string, error) {
	db := database.Pool()

	var task = model.Task{ID: uuid.NewString(), Title: title}
	if err := db.Debug().Create(&task).Error; err != nil {
		return "", err
	}

	return task.ID, nil
}

func UpdateTask(id, title string) error {
	db := database.Pool()

	var task = model.Task{ID: id, Title: title, UpdatedAt: time.Now()}
	if err := db.Debug().Updates(&task).Error; err != nil {
		return err
	}

	return nil
}

func DeleteTask(id string) error {
	db := database.Pool()

	var task = model.Task{ID: id}
	if err := db.Debug().Delete(task).Error; err != nil {
		return err
	}

	return nil
}
