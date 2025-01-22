package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	// Create Echo instance
	e := echo.New()

	// Root endpoint
	// We define a GET endpoint with a dynamic URL parameter.
	// This determines the function that will be run when someone sends a GET request to http://localhost:1212/<name>.
	e.GET("/:name", func(c echo.Context) error {
		name := c.Param("name")                             // Extract the value from the URL parameter.
		return c.String(http.StatusOK, "Hello!, "+name+"!") // Send a response with the extracted name.

	})

	// Start the server and listen on port 1212
	e.Logger.Fatal(e.Start(":1212"))

}
