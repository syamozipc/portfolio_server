package main

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"github.com/syamozipc/web_app/internal/api"
)

func rootGet(c echo.Context) error {
	return c.String(http.StatusOK, "home")
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
	e.GET("/todos", api.ListTodos)
	e.POST("/todos", api.CreateTodo)

	e.Logger.Fatal(e.Start(":8082"))
}
