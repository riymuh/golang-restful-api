package main

import (
	"fmt"
	"net/http"

	"rest-api-golang/handler"
	"rest-api-golang/user"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//connecting to db
	dsn := "u7027509_paibaralek:Paib4ral3k1994__@tcp(45.13.133.19:3306)/u7027509_golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	fmt.Println(db, " ", err)

	db.AutoMigrate(&user.User{})

	testw := "test"
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/user", handler.GetUserListHandler)
	v1.GET("/user/:id", handler.GetUserHandler)
	v1.GET("/user/query", handler.GetUserQueryHandler)
	v1.POST("user/create", handler.PostUserHandler)

	router.GET("/about", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "this is about page!",
			"note":    testw,
		})
	})

	router.Run(":8080")
}
