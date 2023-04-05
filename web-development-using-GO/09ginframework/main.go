//https://www.youtube.com/watch?v=vyfF6lNLE4k&list=PLve39GJ2D71yyECswi0lVaBm_gbnDRR9v&index=10&t=18s

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/aju", GetAju)

	r.POST("/aju", PostAju)

	r.DELETE("/aju", DeleteAju)

	r.Run() // listen and serve on 0.0.0.0:8080   //http://localhost:8080/ping

}

var ajuRatings = map[string]string{ //this is the database we have
	"oneStar": "Beginner star",
}

func GetAju(c *gin.Context) {
	starType := c.Query("type")
	starName, ok := ajuRatings[starType]
	if !ok {
		c.JSON(404, gin.H{
			starType: "",
		})
		return
	}
	c.JSON(200, gin.H{
		starType: starName,
	})
}

func PostAju(c *gin.Context) {
	starType := c.Query("type")
	starName := c.Query("name")

	if len(starType) == 0 || len(starName) == 0 {
		c.JSON(400, gin.H{
			starType: starName,
		})
		return
	}
	//if already resent, then 409 status which is conflict
	if _, ok := ajuRatings[starType]; ok {
		c.JSON(409, gin.H{
			"message": "star already exists",
		})
	} else { //else we can add it to database
		ajuRatings[starType] = starName
		c.JSON(201, gin.H{
			starType: starName,
		})
	}

}

func DeleteAju(c *gin.Context) {
	starType := c.Query("type")
	starName, ok := ajuRatings[starType]
	if !ok {
		c.JSON(404, gin.H{
			starType: "",
		})
		return

	}
	delete(ajuRatings, starType)
	c.JSON(200, gin.H{
		starType: starName,
	})

}

/*

package main
import "github.com/gin-gonic/gin"
func IndexHandler(c *gin.Context){
   c.JSON(200, gin.H{
       "message": "hello world",
   })
}
func main() {
   router := gin.Default()
   router.GET("/", IndexHandler)
   router.Run()
}
*/
