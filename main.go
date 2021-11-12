package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	mpa := map[string]string{
		"key":  "value",
		"key1": "value1",
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		c.Response().Header().Set("Cache-Control", "public, max-age=3600, stale-while-revalidate=30000")
		fmt.Println("new Request")
		return c.JSON(http.StatusOK, mpa)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
