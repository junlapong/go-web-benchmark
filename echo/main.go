package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	// e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}    ${status}    ${method}\t${uri}\n",
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{"message": "pong"})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
