package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	mpa := map[string]string{
		"key":  "value",
		"key1": "value1",
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, mpa)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
