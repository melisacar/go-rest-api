package main

import (
	"net/http"
	"regexp"

	// https://pkg.go.dev/regexp

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// https://pkg.go.dev/github.com/labstack/echo/v4/middleware
)

// Defining the User Struct
type User struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func main() {

	// Echo instance
	e := echo.New()

	// Middleware to log requests
	e.Use(middleware.Logger())

	// Register endpoint
	e.POST("/register", func(c echo.Context) error {

		// Initialize a User struct to bind incoming data
		var user User // Creates a variable user of type User

		// Bind JSON body to the struct
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid request",
			})
		}

		// Validate email format
		if !isValidEmail(user.Email) {
			return c.JSON(http.StatusUnprocessableEntity, map[string]string{
				"error": "Invalid email format",
			})

		}

		// Return success response
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "User registered successfully",
			"user": map[string]string{
				"name":  user.Name,
				"email": user.Email,
			},
		})
	})

	// Start the server and listen on port 1212
	e.Logger.Fatal(e.Start(":1212"))
}

// Email validation function
func isValidEmail(email string) bool {
	// Basic email regex
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}
