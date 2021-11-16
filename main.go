package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net"
	"net/http"
	"os"
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

	e.GET("/uncached", func(c echo.Context) error {
		c.Response().Header().Set("Cache-Control", "no-cache,no-store")
		fmt.Println("Request for uncached data")
		return c.JSON(http.StatusOK, uncachedData)
	})
	printVarnishPodIPs()
	e.Logger.Fatal(e.Start(":1323"))

}

func printVarnishPodIPs() {
	envVariable := os.Getenv("VARNISH_SERVICE_DN")
	fmt.Println("envVariable", envVariable)
	ips, err := net.LookupIP(envVariable)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
		os.Exit(1)
	}
	for _, ip := range ips {
		fmt.Println(ip.To4().String())
	}
}
