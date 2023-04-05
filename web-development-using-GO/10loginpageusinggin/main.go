// not fully completed,, validation not done, just login only
// chatgpt and https://articles.wesionary.team/session-management-in-golang-gin-framework-using-redis-with-e-1f17b6980924
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Render login page
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Login",
		})
	})

	// Handle login form submission
	router.POST("/login", func(c *gin.Context) {
		// Retrieve form data
		// username := c.PostForm("username")
		// password := c.PostForm("password")

		// TODO: Validate username and password

		// Redirect to home page on successful login
		c.Redirect(http.StatusFound, "/home")
	})

	// Render home page
	router.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "welcome.html", gin.H{
			"title": "Home",
		})
	})

	// Serve static files
	router.Static("/static", "./static")

	// Load templates
	router.LoadHTMLGlob("template/*")

	router.GET("/logout", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/")
	})

	// Start server
	router.Run(":8000")
}
