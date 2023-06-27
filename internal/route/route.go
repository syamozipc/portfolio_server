package route

import (
	"github.com/labstack/echo/v4"

	"github.com/syamozipc/web_app/internal/api"
)

func Route(e *echo.Echo) {
	base := e.Group("/api")
	base.GET("", api.Root)

	task := base.Group("/tasks")
	task.GET("/:id", api.GetTask)
	task.GET("", api.ListTasks)
	task.POST("", api.CreateTask)
	task.PATCH("/:id", api.UpdateTask)
}
