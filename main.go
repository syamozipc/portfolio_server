package main

import (
	"flag"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"github.com/syamozipc/web_app/internal/database"
	"github.com/syamozipc/web_app/internal/route"
)

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
	flag.Parse()
	mode := flag.Arg(0)
	e := echo.New()

	switch mode {
	case "server":
		e.Validator = &CustomValidator{validator: validator.New()}
		route.Route(e)
		e.Logger.Fatal(e.Start(":8082"))
	case "migrate":
		err := database.MigrateUp()
		if err != nil {
			e.Logger.Fatal(err)
		}
	}
}
