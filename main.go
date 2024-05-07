package main

import (
	"API/controllers"
	"API/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariabels()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/ping", gin.HandlerFunc(func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ney",
		})
	}))
	

	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostsShow)
	r.DELETE("/posts/:id", controllers.PostsDelete)
	r.Run() // listen and serve on 0.0.0.0:8080

}
