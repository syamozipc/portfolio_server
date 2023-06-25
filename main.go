package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func rootGet(c echo.Context) error {
	return c.String(http.StatusOK, "home")
}

type Person struct {
	Name string
}

func post(c echo.Context) error {
	var p Person
	if err := c.Bind(&p); err != nil {
		return err
	}

	res := struct {
		Message string `json:"message"`
	}{
		Message: fmt.Sprintf("Hello, %s", p.Name),
	}

	return c.JSON(http.StatusCreated, res)
}

func main() {
	e := echo.New()
	e.GET("/", rootGet)
	e.POST("/post", post)

	e.Logger.Fatal(e.Start(":8082"))
}
