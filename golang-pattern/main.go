package main

import (
	"net/http"

	"rest-api-golang/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	testw := "test"
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", handler.RootHandler)
	v1.GET("/user/:id", handler.GetUserHandler)
	v1.GET("/query", handler.GetQueryHandler)
	v1.POST("user/create", handler.PostUserHandler)

	router.GET("/about", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "this is about page!",
			"note":    testw,
		})
	})

	router.Run(":8080")
}
