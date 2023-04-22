//https://betterprogramming.pub/how-to-render-html-pages-with-gin-for-golang-9cb9c8d7e7b6

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	/*
		r := gin.Default()

		// Serve static files from the "template" directory on root URL
		r.Static("/", "./template")

		// Define a route handler for the root route
		r.GET("/", func(c *gin.Context) {
			// Render the "index.html" file using the Bootstrap template
			c.HTML(http.StatusOK, "index.html", gin.H{})
		})

		// Run the server on port 8080
		r.Run()
	*/
	/*
		router := gin.Default()
		router.GET("/hello", func(c *gin.Context) {
			if c != nil {
				c.HTML(http.StatusOK, "template/index.html", gin.H{})
			}
		})

		router.GET("/other", func(c *gin.Context) {
			if c != nil {
				c.String(http.StatusOK, "This is a different routee!")
			}
		})
		router.Run(":8080")
	*/
	/*
		router := gin.Default()
		router.LoadHTMLGlob("template/*.html")

		//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
		router.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "Main website",
			})
		})
		router.Run(":8080")

	*/

	router := gin.Default()
	// router.SetFuncMap(template.FuncMap{
	// 	"upper": strings.ToUpper,
	// })
	router.Static("/assets", "./assets")
	router.LoadHTMLGlob("template/*.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": "This is an index page...",
		})
	})

	router.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "usersignup.html", gin.H{
			"content": "This is an about page...",
		})
	})
	router.Run("localhost:8080")

}

/*
func main() {

	r := gin.Default()

	// Serve static files from the "public" directory
	r.Static("/public", "./public") // works on http://localhost:8080/public/

	// Define a route handler for the root route
	r.GET("/", func(c *gin.Context) {
		// Render the "index.html" file using the Bootstrap template
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	// Run the server on port 8080
	r.Run()

}
*/
