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

/*
func loginHandle(c *gin.Context) {
	// Get form values
	username := c.PostForm("username")
	password := c.PostForm("password")

	if userDataBase[username] == password && password != "" {
		// Get session
		session := sessions.Default(c)

		// Set session value
		session.Set("username", username)

		// Save session
		err := session.Save()
		if err != nil {
			c.String(http.StatusInternalServerError, "Error saving session")
			return
		}

		// Redirect to welcome page
		c.Redirect(http.StatusSeeOther, "/welcome")
	} else {
		// If username and password do not match, render login page with error message
		data := gin.H{
			"error": "Invalid Username and Password Entered",
		}
		c.HTML(http.StatusOK, "index.html", data)
	}
}

*/
