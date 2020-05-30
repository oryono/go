package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oryono/gorm/controllers"
	"github.com/oryono/gorm/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/books", controllers.AllBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.DELETE("/books/:id", controllers.DeleteBook)


	r.Run()
}
