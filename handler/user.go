package handler

import (
	"fmt"
	"net/http"

	"rest-api-golang/user"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetUserListHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World! ini funciton",
	})
}

func GetUserHandler(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello " + id,
	})
}

func GetUserQueryHandler(c *gin.Context) {
	title := c.Query("title")
	genre := c.Query("genre")
	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"genre": genre,
	})
}

func PostUserHandler(c *gin.Context) {
	var userRequest user.UserRequest

	err := c.ShouldBindJSON(&userRequest)

	if err != nil {

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
