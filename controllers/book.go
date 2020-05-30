package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/oryono/gorm/models"
	"net/http"
)

func AllBooks(context *gin.Context)  {
	var books []*models.Book
	models.DB.Find(&books)
	context.JSON(http.StatusOK, gin.H{"data": books})
}

func CreateBook(context *gin.Context)  {

}

func FindBook(context *gin.Context)  {

}

func DeleteBook(context *gin.Context)  {

}
