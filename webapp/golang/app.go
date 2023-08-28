package main

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type User struct {
	ID          int    `db:"id"`
	AccountName string `db:"account_name"`
}

type Pet struct {
	ID     int    `db:"id"`
	UserID int    `db:"user_id"`
	Name   string `db:"name"`
}

type Post struct {
	ID       int       `db:"id"`
	UserID   int       `db:"user_id"`
	Content  string    `db:"content"`
	Comments []Comment `db:"comments"`
}

type Comment struct {
	ID        int       `db:"id"`
	PostID    int       `db:"post_id"`
	UserID    int       `db:"user_id"`
	Comment   string    `db:"comment"`
	CreatedAt time.Time `db:"created_at"`
	User      User
}

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
