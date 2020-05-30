package models

type Author struct {
	ID uint `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}
