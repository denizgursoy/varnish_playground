package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	cachedData := map[string]string{
		"cached": "data",
	}

	uncachedData := map[string]string{
		"uncachedData": "data",
	}

	e := echo.New()

	e.GET("/cached", func(c echo.Context) error {
		c.Response().Header().Set("Cache-Control", "public, max-age=3600, stale-while-revalidate=30000")
		fmt.Println("Request for cached data")
		return c.JSON(http.StatusOK, cachedData)
	})

	e.GET("/uncachedData", func(c echo.Context) error {
		fmt.Println("Request for uncached data")
		return c.JSON(http.StatusOK, uncachedData)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
