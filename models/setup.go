package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB
func ConnectDatabase() {
	database, _ := gorm.Open("sqlite3", "gorm.db")

	database.AutoMigrate(&Book{}, &Author{})
	DB = database

}
