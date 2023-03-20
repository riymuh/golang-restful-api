package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// func main() {
// 	testw := "test"
// 	router := gin.Default()

// 	router.GET("/", rootHandler)
// 	router.GET("/user/:id", getUserHandler)
// 	router.GET("/query", getQueryHandler)
// 	router.POST("user/create", postUserHandler)

// 	router.GET("/about", func(c *gin.Context) {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "this is about page!",
// 			"note":    testw,
// 		})
// 	})

// 	router.Run(":8080")
// }

func main() {
	testw := "test"
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/", rootHandler)
	v1.GET("/user/:id", getUserHandler)
	v1.GET("/query", getQueryHandler)
	v1.POST("user/create", postUserHandler)

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
	StatusName string      `json:"status_name"`
	Phone      json.Number `json:"phone" binding:"required,number"`
}

func postUserHandler(c *gin.Context) {
	var userRequest UserRequest

	err := c.ShouldBindJSON(&userRequest)

	if err != nil {
		//log.Fatal(err)
		// for _, e := range err.(validator.ValidationErrors) {
		// 	errorMessage := fmt.Sprintf("Error on field %s: %v", e.Field(), e.ActualTag())
		// 	c.JSON(http.StatusBadRequest, errorMessage)
		// 	return
		// }

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s: %v", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
		// fmt.Println(err)
		// c.JSON(http.StatusBadRequest, err)
		// return
	}

	c.JSON(http.StatusCreated, gin.H{
		"name":        userRequest.Name,
		"email":       userRequest.Email,
		"status_name": userRequest.StatusName,
		"phone":       userRequest.Phone,
	})
}
