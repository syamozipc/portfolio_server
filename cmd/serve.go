package cmd

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"

	"github.com/syamozipc/web_app/internal/database"
	"github.com/syamozipc/web_app/internal/route"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "boot api server",
	RunE:  serve,
}

func init() {
	rootCmd.AddCommand(serveCmd)
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

func serve(_ *cobra.Command, _ []string) error {
	if err := database.Open(); err != nil {
		return err
	}

	e := echo.New()
	// TODO: urlとか決まったら、CORSWithConfigで詳細を指定
	e.Use(middleware.CORS())

	e.Validator = &CustomValidator{validator: validator.New()}
	route.Route(e)

	return e.Start(":8082")
}
