package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	testw := "test"
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/user/:id", getUserHandler)
	router.GET("/query", getQueryHandler)
	router.POST("user/create", postUserHandler)

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
	title := c.Query("title")
	genre := c.Query("genre")
	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"genre": genre,
	})
}

type UserRequest struct {
	name  string
	email string
	//full_name string
}

func postUserHandler(c *gin.Context) {
	var userRequest UserRequest

	err := c.ShouldBindJSON(&userRequest)

	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"name":  userRequest.name,
		"email": userRequest.email,
		//"full_name": userRequest.full_name,
	})
}
