package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	PURGE = "PURGE"
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

func sendPurgeRequestToAllVarnishPods() {
	for _, ip := range printVarnishPodIPs() {
		client := http.Client{
			Timeout: time.Second * 10,
		}
		client.Do(&http.Request{
			Method: PURGE,
			URL: &url.URL{
				Host: ip.String(),
				Path: "/cached",
			},
		})
	}
}

func getVarnishDomainName() string {
	return os.Getenv("VARNISH_SERVICE_DN")
}

func printVarnishPodIPs() []net.IP {

	ips, err := net.LookupIP(getVarnishDomainName())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
	}
	var ipv4List []net.IP
	for _, ip := range ips {
		if ip.To4() != nil {
			ipv4List = append(ipv4List, ip)
		}
	}
	return ipv4List
}
