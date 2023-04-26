package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Define a route that supports the OPTIONS method
	r.Handle("OPTIONS", "/hello", func(c *gin.Context) {
		// Set the Access-Control-Allow-Origin header to allow cross-origin requests
		c.Header("Access-Control-Allow-Origin", "*")

		// Set the Allow header to list the supported methods
		c.Header("Allow", "GET, POST, OPTIONS")

		// Return a 200 OK response with no body
		c.Status(http.StatusOK)
	})

	// Define a route that returns a "Hello, World!" message
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	// Run the server on port 8080
	r.Run(":8080")
}


/*
In this example, we define a route for the "/hello" endpoint that supports the OPTIONS method. When an OPTIONS request is made to this endpoint, the handler sets the Access-Control-Allow-Origin header to allow cross-origin requests, sets the Allow header to list the supported methods (GET, POST, and OPTIONS), and returns a 200 OK response with no body.

We also define a separate route for the same endpoint that responds with a "Hello, World!" message when a GET request is made. When a GET request is made to the "/hello" endpoint, the handler returns a JSON response with the message.

Finally, we run the Gin server on port 8080. When the server is running, we can make an OPTIONS request to the "/hello" endpoint to retrieve information about the supported methods and headers. For example, we can use the curl command to make an OPTIONS request:


This command should return a 200 OK response with the Allow header set to "GET, POST, OPTIONS".  

(use postman and use OPTIONS request at http://localhost:8080/hello)
