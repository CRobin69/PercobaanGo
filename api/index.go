package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := gin.New()
	server.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "OK",
		})
	})
	server.GET("/hello", func(context *gin.Context) {
		name := context.Query("name")
		if name == "" {
			context.JSON(400, gin.H{
				"message": "name not found",
			})
		} else {
			context.JSON(200, gin.H{
				"data": fmt.Sprintf("Hello %s!", name),
			})
		}
	})
	server.GET("/user/:id", func(context *gin.Context) {
		context.JSON(400, gin.H{
			"data": gin.H{
				"id": context.Param("id"),
			},
		})
	})
	server.ServeHTTP(w, r)
}
