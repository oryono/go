package models

import (
	_ "github.com/jinzhu/gorm"
)

type Book struct {
	ID uint `json:"id" gorm:"primary_key"`
	Title string `json:"title"`
	AuthorId string `json:"author"`
}
