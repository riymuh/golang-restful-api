package main

import (
	"fmt"
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
	Name       string `json:"name" binding:"required"`
	Email      string
	StatusName string `json:"status_name"`
}

func postUserHandler(c *gin.Context) {
	var userRequest UserRequest

	err := c.ShouldBindJSON(&userRequest)

	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"name":        userRequest.Name,
		"email":       userRequest.Email,
		"status_name": userRequest.StatusName,
	})
}
