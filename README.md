# Building RESTful APIs with Go

This guide demonstrates how to build a series of RESTful APIs using the [Echo](https://echo.labstack.com/) web framework in Go. We cover:

1. Setting up a simple Echo web server.
2. Creating a server with dynamic URL parameters.
3. Implementing a user registration API with validation.

---

## Requirements

- [Go](https://go.dev/dl/) (version 1.16 or higher recommended)
- A code editor (e.g., [VS Code](https://code.visualstudio.com/))
- [Postman](https://www.postman.com/) or a similar tool for testing APIs
- Internet connection to download dependencies

---

## 1. Simple Echo Web Server in Go

This example demonstrates how to set up a basic web server using the [Echo](https://echo.labstack.com/) framework in Go. The server listens on port `1212` and serves a simple `GET` endpoint that returns a "Hello!" response.

### Installation

1. Create a new directory for the project:

   ```bash
   mkdir echo-web-server
   cd echo-web-server
   ```

2. Initialize the project and install Echo:

    ```bash
    go mod init echo-web-server
    go get github.com/labstack/echo/v4
    ```

3. Create a file named `main.go` and paste the following code:

    ```go
    package main

    import (
    "net/http"

    "github.com/labstack/echo/v4"
    )

    func main() {

    // Create Echo instance
    e := echo.New()

    // Root endpoint
    // We define a GET endpoint.
    // This determines the function that will be run when someone sends a GET request to http://localhost:8080.
    e.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK, "Hello!") // We send a text (string) response.

    })

    // Start the server and listen on port 1212
    e.Logger.Fatal(e.Start(":1212"))

    }
    ```

4. Clean up dependencies

    ```bash
    go mod tidy
    ```

### Running the Server

1. Run the Server with:

    ```bash
    go run main.go
    ```

2. The server will start listening on port 1212. Open your browser or use a tool like Postman to make a GET request to:

    ```bash
    http://localhost:1212/
    ```

3. You should see the response:

    ```bash
    Hello!
    ```

---

## 2. Go Echo Server with Dynamic URL Parameter

This example demonstrates how to use dynamic URL parameters to personalize responses.

### Installation-2

Update the `main.go` file with the following code:

```go
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
```

### Testing the Server

1. Start the server:

    ```bash
    go run main.go
    ```

2. Open your browser or Postman and navigate to:

    ```bash
    http://localhost:1212/YourName
    ```

3. Replace `YourName` with any name to receive a personalized greeting.

---

## 3. User Registration API

This example is a simple Go-based REST API for user registration. It accepts user details, validates the input, and returns a response.

### Features

- **POST** `/register` endpoint for user registration.
- Validates the incoming JSON request:
  - Ensures `name`, `email`, and `password` fields are provided.
  - Checks if the `email` is in a valid format.
- Returns appropriate error responses for invalid input.
- Logs requests using middleware.

### Installation-3

Update the `main.go` file with the following code:

```go
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

```

### Testing the API

1. Start the server:

    ```bash
    go run main.go
    ```

2. Use Postman to send a `POST` request to:

    ```bash
    http://localhost:1212/register
    ```

3. Example request body:

    ```json
    {
        "name": "Melisa Acar",
        "email": "melisacar@example.com",
        "password": "secret"
    }
    ```

4. Responses:

    - Success (200 OK):

    ```json
    {
        "message": "User registered successfully",
        "user": {
            "name": "Melisa Acar",
            "email": "melisacar@example.com"
        }
    }
    ```

    - Invalid Email (422 Unprocessable Entity):

    ```json
    {
        "error": "Invalid email format"
    }
    ```

    - Invalid Request (400 Bad Request):

    ```json
    {
        "error": "Invalid request"
    }
    ```

### Steps for Testing with Postman

1. Open Postman and create a new `POST` request.
2. Set the URL to: `http://localhost:1212/register`.
3. Go to the Body tab, select `raw`, and set the format to `JSON`.
4. Enter the request body given above.
5. Click Send.

---

## Notes

- Unused Imports and Variables: Go does not allow unused imports or variables. Tools like `gofmt` and `goimports` help maintain code standards.
- Dependency Management: Use `go mod tidy` to clean up dependencies.