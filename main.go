package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	testw := "test"
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/user/:id", getUserHandler)
	router.GET("/query", getQueryHandler)

	router.GET("/about", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "this is about page!",
			"note":    testw,
		})
	})

	router.Run(":8080")
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World! ini funciton",
	})
}

func getUserHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello " + id,
	})
}

func getQueryHandler(c *gin.Context) {
	id := c.Query("title")
	c.JSON(http.StatusOK, gin.H{
		"title_test": id,
	})
}
