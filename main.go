package main

import (
	"TestTask/posts"
	"TestTask/storage"
	"github.com/labstack/echo/v4"
)

func main() {

	// Set up database connection
	storage.DBConn()

	// Create Echo instance
	e := echo.New()

	// Route for creating a new check
	e.POST("/hash", posts.Hash)
	e.POST("/records", posts.Records)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
