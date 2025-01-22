# Builidng RESTful APIs with Go

## Simple Echo Web Server in Go

This part demonstrates how to set up a basic web server using the [Echo](https://echo.labstack.com/) framework in Go. The server listens on port `1212` and serves a simple `GET` endpoint that returns a "Hello!" response.

### Requirements

- [Go](https://go.dev/dl/) (version 1.16 or higher is recommended)
- Internet connection to download the Echo package (if not already installed)

### Installation

1. Clone the repository or create a new directory for the project:

```bash
mkdir echo-web-server
cd echo-web-server
```

2. Initialize the project and install dependencies:

```bash
go mod init echo-web-server
go get github.com/labstack/echo/v4
```

Create a file named main.go and paste the following code:

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

3. Clean up dependencies

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

## Caution

In Go, unused imports or variables are not allowed, and the compiler enforces this rule. When you save a file in VS Code, the Go extension often triggers goimports or gofmt, which automatically removes unused imports or formats your code according to Go's standards.

---

## Go Echo Server with Dynamic URL Parameter

This second Go application uses the Echo framework to create a simple web server with a dynamic URL parameter.

Example:

- **Request**: `http://localhost:1212/Melisa`
- **Response**: `Hello, Melisa!`

### Installation

Paste the following code to the main.go file:

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

### Setup and Run

1. Clone the repository or create a Go project with the following code.

2. Install the necessary dependencies:

```bash
go get github.com/labstack/echo/v4
```

3. Run the Go application:

```bash
go run main.go
```

4. Open your browser and navigate to `http://localhost:1212/yourname`, replacing `yourname` with any name you like. The server will respond with a personalized greeting.

---