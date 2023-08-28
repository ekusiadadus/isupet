package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// if APP_PORT is not set, use default port 1323
	APP_PORT := os.Getenv("APP_PORT")
	if APP_PORT == "" {
		APP_PORT = "1323"
	}
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":" + APP_PORT))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
