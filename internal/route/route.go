package route

import (
	"github.com/labstack/echo/v4"

	"github.com/syamozipc/web_app/internal/api"
)

func Route(e *echo.Echo) {
	e.GET("/", api.Home)
	e.GET("/tasks", api.ListTasks)
	e.POST("/tasks", api.CreateTask)
}
